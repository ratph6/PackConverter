package converter

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

type Converter struct {
	InputPath  string
	OutputPath string
	OnLog      func(string)
	OnProgress func(float64)
}

type Stats struct {
	TotalFiles    int
	RenamedFiles  int
	ModifiedFiles int
	SplitFiles    int
	Errors        int
}

func (c *Converter) log(msg string) {
	if c.OnLog != nil {
		c.OnLog(msg)
	}
}

func (c *Converter) progress(p float64) {
	if c.OnProgress != nil {
		c.OnProgress(p)
	}
}

func DetectSourceVersion(reader *zip.ReadCloser) string {
	for _, file := range reader.File {
		if path.Base(file.Name) == "pack.mcmeta" {
			rc, err := file.Open()
			if err != nil {
				continue
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				continue
			}
			format := extractPackFormat(data)
			switch {
			case format <= 1:
				return "1.6.1-1.8.9"
			case format == 2:
				return "1.9-1.10.2"
			case format == 3:
				return "1.11-1.12.2"
			case format == 4:
				return "1.13-1.14.4"
			case format == 5:
				return "1.15-1.16.1"
			case format <= 7:
				return "1.16.2-1.17.1"
			case format <= 9:
				return "1.18-1.19.2"
			case format <= 15:
				return "1.19.3-1.20.1"
			case format <= 34:
				return "1.20.2-1.21.1"
			default:
				return fmt.Sprintf("1.21+ (format %d)", format)
			}
		}
	}
	for _, file := range reader.File {
		if strings.Contains(file.Name, "textures/blocks/") || strings.Contains(file.Name, "textures/items/") {
			return "1.6.1-1.8.9 (detected from directory structure)"
		}
	}
	return "Unknown"
}

type outputFile struct {
	path string
	data []byte
}

type processResult struct {
	files    []outputFile
	logs     []string
	err      error
	renamed  int
	modified int
	split    int
}

type fileEntry struct {
	name  string
	data  []byte
	isDir bool
}

func (c *Converter) Convert() (*Stats, error) {
	reader, err := zip.OpenReader(c.InputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input zip: %w", err)
	}
	defer reader.Close()

	sourceVersion := DetectSourceVersion(reader)
	c.log(fmt.Sprintf("Detected source version: %s", sourceVersion))
	c.log(fmt.Sprintf("Target: Minecraft 1.21.10 (pack_format %d)", TargetPackFormat))

	stats := &Stats{TotalFiles: len(reader.File)}
	c.log(fmt.Sprintf("Processing %d files...\n", stats.TotalFiles))

	// Read all file data from zip sequentially (zip format requires this).
	entries := make([]fileEntry, len(reader.File))
	for i, file := range reader.File {
		if file.FileInfo().IsDir() {
			entries[i] = fileEntry{name: file.Name, isDir: true}
			continue
		}
		rc, err := file.Open()
		if err != nil {
			entries[i] = fileEntry{name: file.Name}
			continue
		}
		size := file.UncompressedSize64
		if size == 0 || size > 256<<20 {
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				entries[i] = fileEntry{name: file.Name}
				continue
			}
			entries[i] = fileEntry{name: file.Name, data: data}
		} else {
			data := make([]byte, size)
			_, err := io.ReadFull(rc, data)
			rc.Close()
			if err != nil {
				entries[i] = fileEntry{name: file.Name}
				continue
			}
			entries[i] = fileEntry{name: file.Name, data: data}
		}
	}

	// Process files in parallel.
	results := make([]processResult, len(entries))
	numWorkers := runtime.NumCPU()
	if numWorkers < 2 {
		numWorkers = 2
	}

	var processed int64
	total := int64(stats.TotalFiles)

	jobs := make(chan int, numWorkers*2)
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				results[idx] = processFileData(entries[idx].name, entries[idx].data)
				done := atomic.AddInt64(&processed, 1)
				c.progress(float64(done) / float64(total) * 0.9) // 0-90%
			}
		}()
	}

	for i := range entries {
		if entries[i].isDir || entries[i].data == nil {
			atomic.AddInt64(&processed, 1)
			continue
		}
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// Write results to output zip, streaming to disk.
	c.log("Writing output zip...")
	outFile, err := os.Create(c.OutputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create output file: %w", err)
	}

	bufWriter := bufio.NewWriterSize(outFile, 256<<10)
	writer := zip.NewWriter(bufWriter)

	writtenFiles := make(map[string]bool)
	for i, r := range results {
		for _, msg := range r.logs {
			c.log(msg)
		}
		if r.err != nil {
			c.log(fmt.Sprintf("ERROR: %s: %v", entries[i].name, r.err))
			stats.Errors++
		}

		stats.RenamedFiles += r.renamed
		stats.ModifiedFiles += r.modified
		stats.SplitFiles += r.split

		for _, nf := range r.files {
			if writtenFiles[nf.path] {
				continue
			}
			header := &zip.FileHeader{
				Name:   nf.path,
				Method: compressionMethod(nf.path),
			}
			w, err := writer.CreateHeader(header)
			if err != nil {
				stats.Errors++
				continue
			}
			if _, err := w.Write(nf.data); err != nil {
				stats.Errors++
				continue
			}
			writtenFiles[nf.path] = true
		}

		c.progress(0.9 + float64(i+1)/float64(len(results))*0.1) // 90-100%
	}

	if err := writer.Close(); err != nil {
		outFile.Close()
		os.Remove(c.OutputPath)
		return nil, fmt.Errorf("failed to finalize zip: %w", err)
	}
	if err := bufWriter.Flush(); err != nil {
		outFile.Close()
		os.Remove(c.OutputPath)
		return nil, fmt.Errorf("failed to flush output: %w", err)
	}
	if err := outFile.Close(); err != nil {
		os.Remove(c.OutputPath)
		return nil, fmt.Errorf("failed to close output file: %w", err)
	}

	c.log(fmt.Sprintf("\nDone! %d files processed, %d renamed, %d modified, %d split, %d errors",
		stats.TotalFiles, stats.RenamedFiles, stats.ModifiedFiles, stats.SplitFiles, stats.Errors))

	return stats, nil
}

// compressionMethod returns Store for pre-compressed formats, Deflate otherwise.
func compressionMethod(filePath string) uint16 {
	switch strings.ToLower(path.Ext(filePath)) {
	case ".png", ".jpg", ".jpeg", ".ogg", ".mp3", ".zip", ".gz", ".webp":
		return zip.Store
	default:
		return zip.Deflate
	}
}

// processFileData handles a single file. Safe for concurrent use.
func processFileData(filePath string, data []byte) processResult {
	var r processResult
	newPath := filePath
	renamed := false
	modified := false

	if path.Base(filePath) == "pack.mcmeta" {
		newData, err := ConvertPackMeta(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not parse pack.mcmeta, copying as-is: %v", err))
		} else {
			data = newData
			modified = true
			r.logs = append(r.logs, fmt.Sprintf("Updated pack.mcmeta: pack_format → %d", TargetPackFormat))
		}
	}

	if strings.HasSuffix(filePath, ".lang") && strings.Contains(filePath, "/lang/") {
		newData, newFilename, err := ConvertLangToJSON(data, path.Base(filePath))
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not convert lang file %s: %v", filePath, err))
		} else {
			data = newData
			dir := path.Dir(filePath)
			newPath = dir + "/" + newFilename
			modified = true
			renamed = true
			r.logs = append(r.logs, fmt.Sprintf("Converted lang: %s → %s", path.Base(filePath), newFilename))
		}
	}

	if path.Base(filePath) == "sounds.json" {
		newData, changed, err := ConvertSoundsJSON(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not convert sounds.json: %v", err))
		} else if changed {
			data = newData
			modified = true
			r.logs = append(r.logs, "Updated sounds.json: remapped sound events")
		}
	}

	if strings.Contains(filePath, "/models/") && strings.HasSuffix(filePath, ".json") {
		newData, changed, err := ConvertModelJSON(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not convert model %s: %v", filePath, err))
		} else if changed {
			data = newData
			modified = true
		}
	}

	if strings.Contains(filePath, "/blockstates/") && strings.HasSuffix(filePath, ".json") {
		newData, changed, err := ConvertBlockstateJSON(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not convert blockstate %s: %v", filePath, err))
		} else if changed {
			data = newData
			modified = true
		}
	}

	if isPaintingAtlas(filePath) {
		results, err := SplitPaintingAtlas(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not split painting atlas: %v", err))
		} else {
			dir := path.Dir(filePath)
			for name, imgData := range results {
				r.files = append(r.files, outputFile{dir + "/" + name, imgData})
			}
			r.logs = append(r.logs, fmt.Sprintf("Split painting atlas → %d individual paintings", len(results)))
			r.split = len(results)
			return r
		}
	}

	if isWidgetsAtlas(filePath) {
		results, err := SplitWidgetAtlas(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not split widgets atlas: %v", err))
		} else if len(results) > 0 {
			if idx := strings.Index(filePath, "textures/gui/"); idx >= 0 {
				guiBase := filePath[:idx+len("textures/gui/")]
				for name, imgData := range results {
					r.files = append(r.files, outputFile{guiBase + "sprites/" + name, imgData})
				}
				r.logs = append(r.logs, fmt.Sprintf("Split widgets.png → %d GUI sprites", len(results)))
				r.split = len(results)
				r.files = append(r.files, outputFile{newPath, data}) // keep original
				return r
			}
		}
	}

	if isIconsAtlas(filePath) {
		results, err := SplitIconsAtlas(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not split icons atlas: %v", err))
		} else if len(results) > 0 {
			if idx := strings.Index(filePath, "textures/gui/"); idx >= 0 {
				guiBase := filePath[:idx+len("textures/gui/")]
				for name, imgData := range results {
					r.files = append(r.files, outputFile{guiBase + "sprites/" + name, imgData})
				}
				r.logs = append(r.logs, fmt.Sprintf("Split icons.png → %d HUD sprites", len(results)))
				r.split = len(results)
				r.files = append(r.files, outputFile{newPath, data})
				return r
			}
		}
	}

	if isDoubleChestTexture(filePath) {
		left, right, err := SplitDoubleChest(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not split double chest texture: %v", err))
		} else {
			dir := path.Dir(filePath)
			baseName := strings.TrimSuffix(path.Base(filePath), ".png")
			baseName = strings.TrimSuffix(baseName, "_double")
			r.logs = append(r.logs, fmt.Sprintf("Split double chest: %s → left + right", path.Base(filePath)))
			r.split = 2
			r.files = []outputFile{
				{dir + "/" + baseName + "_left.png", left},
				{dir + "/" + baseName + "_right.png", right},
			}
			return r
		}
	}

	if isPanoramaFace(filePath) && strings.HasSuffix(filePath, ".png") {
		newData, err := squarePanoramaFace(data)
		if err != nil {
			r.logs = append(r.logs, fmt.Sprintf("WARNING: Could not fix panorama face %s: %v", filePath, err))
		} else if newData != nil {
			data = newData
			modified = true
			r.logs = append(r.logs, fmt.Sprintf("Fixed panorama face: %s (made square)", path.Base(filePath)))
		}
	}

	if !renamed {
		remapped, wasRenamed := RemapPath(newPath)
		if wasRenamed {
			newPath = remapped
			renamed = true
		}
	}

	if renamed {
		r.renamed = 1
	}
	if modified {
		r.modified = 1
	}

	// Old MC had one enchant glint texture; 1.21 splits it into item + entity.
	if isEnchantGlint(newPath) {
		entityPath := strings.Replace(newPath, "enchanted_glint_item", "enchanted_glint_entity", 1)
		r.files = []outputFile{{newPath, data}, {entityPath, data}}
		r.logs = append(r.logs, "Duplicated enchant glint → item + entity variants")
		return r
	}

	r.files = []outputFile{{newPath, data}}
	return r
}

func isPaintingAtlas(p string) bool {
	base := strings.TrimSuffix(path.Base(p), path.Ext(p))
	return base == PaintingAtlasName &&
		(strings.Contains(p, "textures/painting/") || strings.Contains(p, "textures/art/"))
}

func isWidgetsAtlas(p string) bool {
	return path.Base(p) == "widgets.png" && strings.Contains(p, "textures/gui/")
}

func isIconsAtlas(p string) bool {
	return path.Base(p) == "icons.png" && strings.Contains(p, "textures/gui/")
}

func isEnchantGlint(p string) bool {
	base := path.Base(p)
	return base == "enchanted_glint_item.png" || base == "enchanted_glint_item.png.mcmeta"
}

func isDoubleChestTexture(p string) bool {
	return strings.Contains(p, "textures/entity/chest/") && strings.HasSuffix(path.Base(p), "_double.png")
}

func extractPackFormat(data []byte) int {
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})
	cleaned := stripControlCharsBytes(data)
	var meta PackMeta
	if err := json.Unmarshal(cleaned, &meta); err != nil {
		return -1
	}
	return meta.Pack.PackFormat
}

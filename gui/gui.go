package gui

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"

	"texture-pack-converter/converter"
)

func Run() {
	a := app.New()
	w := a.NewWindow("MC Texture Pack Converter")
	w.Resize(fyne.NewSize(720, 540))

	var inputPath, outputPath string

	// Path labels
	inputPathLabel := widget.NewLabel("No file selected")
	inputPathLabel.Wrapping = fyne.TextTruncate
	outputPathLabel := widget.NewLabel("Auto-set when you pick an input")
	outputPathLabel.Wrapping = fyne.TextTruncate

	// Status line (shows live stats during/after conversion)
	statusLabel := widget.NewLabel("")

	// Progress bar
	progress := widget.NewProgressBar()
	progress.Hide()

	// Log area
	logText := widget.NewMultiLineEntry()
	logText.Wrapping = fyne.TextWrapWord
	logText.Disable()
	logScroll := container.NewScroll(logText)
	logScroll.SetMinSize(fyne.NewSize(0, 260))

	// Buttons
	convertBtn := widget.NewButton("Convert", nil)
	convertBtn.Importance = widget.HighImportance
	convertBtn.Disable()

	updateConvertBtn := func() {
		if inputPath != "" && outputPath != "" {
			convertBtn.Enable()
		} else {
			convertBtn.Disable()
		}
	}

	inputBtn := widget.NewButton("Browse...", nil)
	outputBtn := widget.NewButton("Change...", nil)

	// Native file open dialog — no file created, just returns a path
	inputBtn.OnTapped = func() {
		go func() {
			file, err := zenity.SelectFile(
				zenity.Title("Select Input Texture Pack"),
				zenity.FileFilters{
					{Name: "ZIP files", Patterns: []string{"*.zip"}, CaseFold: true},
				},
			)
			if err != nil || file == "" {
				return
			}
			inputPath = file
			inputPathLabel.SetText(file)

			// Auto-generate output path next to the input
			dir := filepath.Dir(file)
			base := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			outputPath = filepath.Join(dir, base+"_1.21.zip")
			outputPathLabel.SetText(outputPath)
			updateConvertBtn()
		}()
	}

	// Native file save dialog — just returns a path, doesn't create the file
	outputBtn.OnTapped = func() {
		go func() {
			opts := []zenity.Option{
				zenity.Title("Save Converted Pack As"),
				zenity.ConfirmOverwrite(),
				zenity.FileFilters{
					{Name: "ZIP files", Patterns: []string{"*.zip"}, CaseFold: true},
				},
			}
			if outputPath != "" {
				opts = append(opts, zenity.Filename(outputPath))
			}
			file, err := zenity.SelectFileSave(opts...)
			if err != nil || file == "" {
				return
			}
			if !strings.HasSuffix(strings.ToLower(file), ".zip") {
				file += ".zip"
			}
			outputPath = file
			outputPathLabel.SetText(file)
			updateConvertBtn()
		}()
	}

	convertBtn.OnTapped = func() {
		convertBtn.Disable()
		inputBtn.Disable()
		outputBtn.Disable()
		progress.Show()
		progress.SetValue(0)
		logText.SetText("")
		statusLabel.SetText("Processing...")

		go func() {
			var mu sync.Mutex
			var logLines []string
			lastFlush := time.Now()

			flushLog := func() {
				text := strings.Join(logLines, "\n")
				logText.SetText(text)
				logScroll.ScrollToBottom()
				lastFlush = time.Now()
			}

			conv := &converter.Converter{
				InputPath:  inputPath,
				OutputPath: outputPath,
				OnLog: func(msg string) {
					mu.Lock()
					defer mu.Unlock()
					logLines = append(logLines, msg)
					// Throttle UI updates to avoid lag on large packs
					if time.Since(lastFlush) > 100*time.Millisecond {
						flushLog()
					}
				},
				OnProgress: func(p float64) {
					progress.SetValue(p)
				},
			}

			start := time.Now()
			stats, err := conv.Convert()
			elapsed := time.Since(start)

			// Final log flush
			mu.Lock()
			flushLog()
			mu.Unlock()

			if err != nil {
				statusLabel.SetText("Conversion failed!")
				progress.Hide()
				dialog.ShowError(err, w)
			} else {
				statusLabel.SetText(fmt.Sprintf(
					"Done in %v  |  %d files  |  %d renamed  |  %d modified  |  %d split  |  %d errors",
					elapsed.Round(time.Millisecond),
					stats.TotalFiles, stats.RenamedFiles, stats.ModifiedFiles,
					stats.SplitFiles, stats.Errors,
				))
				progress.Hide()
				dialog.ShowInformation("Done",
					fmt.Sprintf("Converted in %v\nOutput: %s", elapsed.Round(time.Millisecond), outputPath), w)
			}
			convertBtn.Enable()
			inputBtn.Enable()
			outputBtn.Enable()
		}()
	}

	// Layout
	inputRow := container.NewBorder(nil, nil,
		widget.NewLabelWithStyle("Input:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		inputBtn,
		inputPathLabel,
	)
	outputRow := container.NewBorder(nil, nil,
		widget.NewLabelWithStyle("Output:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		outputBtn,
		outputPathLabel,
	)

	header := container.NewVBox(
		widget.NewLabelWithStyle("MC Texture Pack Converter",
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("1.7 / 1.8.9 → 1.21.10",
			fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewSeparator(),
		inputRow,
		outputRow,
		widget.NewSeparator(),
		container.NewHBox(layout.NewSpacer(), convertBtn, layout.NewSpacer()),
		progress,
		statusLabel,
	)

	content := container.NewBorder(header, nil, nil, nil, logScroll)
	w.SetContent(content)
	w.ShowAndRun()
}

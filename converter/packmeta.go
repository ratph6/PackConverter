package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const TargetPackFormat = 69 // Minecraft 1.21.10

type PackMeta struct {
	Pack PackInfo `json:"pack"`
}

type SupportedFormats struct {
	MinInclusive int `json:"min_inclusive"`
	MaxInclusive int `json:"max_inclusive"`
}

type PackInfo struct {
	PackFormat       int              `json:"pack_format"`
	SupportedFormats SupportedFormats `json:"supported_formats"`
	Description      any              `json:"description"`
}

func ConvertPackMeta(data []byte) ([]byte, error) {
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF}) // strip BOM
	cleaned := stripControlCharsBytes(data)

	var meta PackMeta
	if err := json.Unmarshal(cleaned, &meta); err != nil {
		return fallbackPackMeta(cleaned)
	}

	meta.Pack.PackFormat = TargetPackFormat
	meta.Pack.SupportedFormats = SupportedFormats{
		MinInclusive: 18,
		MaxInclusive: TargetPackFormat,
	}
	meta.Pack.Description = sanitizeDescription(meta.Pack.Description)

	out, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return fallbackPackMeta(cleaned)
	}

	return stripControlCharsBytes(out), nil
}

func fallbackPackMeta(data []byte) ([]byte, error) {
	re := regexp.MustCompile(`"pack_format"\s*:\s*\d+`)
	replacement := fmt.Sprintf(`"pack_format": %d,
    "supported_formats": {
      "min_inclusive": 18,
      "max_inclusive": %d
    }`, TargetPackFormat, TargetPackFormat)
	replaced := re.ReplaceAll(data, []byte(replacement))
	if !bytes.Equal(replaced, data) {
		return stripControlCharsBytes(replaced), nil
	}

	// Last resort: minimal valid pack.mcmeta.
	fallback := fmt.Sprintf(`{
  "pack": {
    "pack_format": %d,
    "supported_formats": {
      "min_inclusive": 18,
      "max_inclusive": %d
    },
    "description": "Converted resource pack"
  }
}`, TargetPackFormat, TargetPackFormat)
	return []byte(fallback), nil
}

func sanitizeDescription(desc any) any {
	switch v := desc.(type) {
	case string:
		return stripControlChars(v)
	case map[string]any:
		for key, val := range v {
			v[key] = sanitizeDescription(val)
		}
		return v
	case []any:
		for i, val := range v {
			v[i] = sanitizeDescription(val)
		}
		return v
	default:
		return desc
	}
}

func stripControlChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\r' && r != '\t' {
			return -1
		}
		return r
	}, s)
}

// stripControlCharsBytes removes control characters from JSON byte data while
// respecting string boundaries and escape sequences.
func stripControlCharsBytes(data []byte) []byte {
	var buf bytes.Buffer
	buf.Grow(len(data))
	inString := false
	escaped := false
	for _, b := range data {
		if escaped {
			buf.WriteByte(b)
			escaped = false
			continue
		}
		if b == '\\' && inString {
			buf.WriteByte(b)
			escaped = true
			continue
		}
		if b == '"' {
			inString = !inString
		}
		if inString && b < 0x20 && b != '\t' && b != '\n' && b != '\r' {
			continue
		}
		buf.WriteByte(b)
	}
	return buf.Bytes()
}

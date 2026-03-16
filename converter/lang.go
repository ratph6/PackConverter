package converter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ConvertLangToJSON converts a .lang file (key=value) to .json format.
// Also lowercases the filename (en_US.lang → en_us.json).
func ConvertLangToJSON(data []byte, filename string) ([]byte, string, error) {
	translations := make(map[string]string)

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx < 0 {
			continue
		}
		translations[line[:idx]] = line[idx+1:]
	}

	if err := scanner.Err(); err != nil {
		return nil, "", fmt.Errorf("failed to read lang file: %w", err)
	}

	out, err := json.MarshalIndent(translations, "", "  ")
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal lang json: %w", err)
	}

	newFilename := strings.ToLower(strings.TrimSuffix(filename, ".lang")) + ".json"
	return out, newFilename, nil
}

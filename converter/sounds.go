package converter

import (
	"encoding/json"
	"fmt"
)

// ConvertSoundsJSON remaps sound event names from 1.8.9 to 1.21 format.
func ConvertSoundsJSON(data []byte) ([]byte, bool, error) {
	var sounds map[string]any
	if err := json.Unmarshal(data, &sounds); err != nil {
		return data, false, nil
	}

	changed := false
	newSounds := make(map[string]any, len(sounds))

	for eventName, eventData := range sounds {
		newName := eventName
		if mapped, ok := SoundEventRenames[eventName]; ok {
			newName = mapped
			changed = true
		}
		newSounds[newName] = eventData
	}

	if !changed {
		return data, false, nil
	}

	out, err := json.MarshalIndent(newSounds, "", "  ")
	if err != nil {
		return nil, false, fmt.Errorf("failed to marshal sounds.json: %w", err)
	}
	return out, true, nil
}

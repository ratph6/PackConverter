package converter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConvertModelJSON updates texture and parent references in model JSON files.
func ConvertModelJSON(data []byte) ([]byte, bool, error) {
	var model map[string]any
	if err := json.Unmarshal(data, &model); err != nil {
		return data, false, nil
	}

	changed := false

	if parent, ok := model["parent"].(string); ok {
		newParent := remapModelReference(parent)
		if newParent != parent {
			model["parent"] = newParent
			changed = true
		}
	}

	if textures, ok := model["textures"].(map[string]any); ok {
		for key, val := range textures {
			if ref, ok := val.(string); ok {
				newRef := RemapTextureReference(ref)
				if newRef != ref {
					textures[key] = newRef
					changed = true
				}
			}
		}
	}

	if !changed {
		return data, false, nil
	}

	out, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return nil, false, fmt.Errorf("failed to marshal model: %w", err)
	}
	return out, true, nil
}

// ConvertBlockstateJSON updates model references in blockstate files.
func ConvertBlockstateJSON(data []byte) ([]byte, bool, error) {
	var blockstate map[string]any
	if err := json.Unmarshal(data, &blockstate); err != nil {
		return data, false, nil
	}

	changed := false

	if variants, ok := blockstate["variants"].(map[string]any); ok {
		for _, val := range variants {
			if updateModelRefs(val) {
				changed = true
			}
		}
	}

	if multipart, ok := blockstate["multipart"].([]any); ok {
		for _, part := range multipart {
			if partMap, ok := part.(map[string]any); ok {
				if apply, ok := partMap["apply"]; ok {
					if updateModelRefs(apply) {
						changed = true
					}
				}
			}
		}
	}

	if !changed {
		return data, false, nil
	}

	out, err := json.MarshalIndent(blockstate, "", "  ")
	if err != nil {
		return nil, false, fmt.Errorf("failed to marshal blockstate: %w", err)
	}
	return out, true, nil
}

func updateModelRefs(val any) bool {
	changed := false
	switch v := val.(type) {
	case map[string]any:
		if model, ok := v["model"].(string); ok {
			newModel := remapModelReference(model)
			if newModel != model {
				v["model"] = newModel
				changed = true
			}
		}
	case []any:
		for _, item := range v {
			if updateModelRefs(item) {
				changed = true
			}
		}
	}
	return changed
}

func remapModelReference(ref string) string {
	cleanRef := strings.TrimPrefix(ref, "minecraft:")

	if strings.HasPrefix(cleanRef, "block/") {
		name := cleanRef[len("block/"):]
		if newName, ok := BlockTextureRenames[name]; ok {
			return "minecraft:block/" + newName
		}
		return "minecraft:" + cleanRef
	}

	if strings.HasPrefix(cleanRef, "item/") {
		name := cleanRef[len("item/"):]
		if newName, ok := ItemTextureRenames[name]; ok {
			return "minecraft:item/" + newName
		}
		return "minecraft:" + cleanRef
	}

	// Old format without namespace prefix: assume block.
	if !strings.Contains(cleanRef, "/") && !strings.Contains(cleanRef, ":") {
		if newName, ok := BlockTextureRenames[cleanRef]; ok {
			return "minecraft:block/" + newName
		}
		return "minecraft:block/" + cleanRef
	}

	return ref
}

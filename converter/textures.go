package converter

import (
	"path"
	"strings"
)

// RemapPath converts a file path from 1.7/1.8.9 to 1.21.10 naming.
func RemapPath(filePath string) (string, bool) {
	newPath := filePath
	renamed := false

	// Directory renames: blocks/ → block/, items/ → item/
	for oldDir, newDir := range DirectoryRenames {
		if idx := strings.Index(newPath, oldDir); idx >= 0 {
			newPath = newPath[:idx] + newDir + newPath[idx+len(oldDir):]
			renamed = true
			break
		}
	}

	dir := path.Dir(newPath)
	base := path.Base(newPath)
	ext := path.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// .png.mcmeta: strip both extensions for lookup, re-add after
	if ext == ".mcmeta" && strings.HasSuffix(name, ".png") {
		name = strings.TrimSuffix(name, ".png")
		ext = ".png.mcmeta"
	}

	if isBlockTextureDir(dir) {
		if newName, ok := BlockTextureRenames[name]; ok {
			return dir + "/" + newName + ext, true
		}
	}

	if isItemTextureDir(dir) {
		if newName, ok := ItemTextureRenames[name]; ok {
			return dir + "/" + newName + ext, true
		}
	}

	if idx := strings.Index(newPath, "textures/entity/"); idx >= 0 {
		prefix := newPath[:idx+len("textures/entity/")]
		entityRel := newPath[idx+len("textures/entity/"):]
		entityName := strings.TrimSuffix(entityRel, ext)
		if newName, ok := EntityTextureRenames[entityName]; ok {
			return prefix + newName + ext, true
		}
	}

	if idx := strings.Index(newPath, "textures/models/armor/"); idx >= 0 {
		prefix := newPath[:idx+len("textures/")]
		armorName := strings.TrimSuffix(newPath[idx+len("textures/"):], ext)
		if newName, ok := ArmorLayerRenames[armorName]; ok {
			return prefix + newName + ext, true
		}
	}

	if idx := strings.Index(newPath, "textures/entity/horse/armor/"); idx >= 0 {
		prefix := newPath[:idx+len("textures/")]
		armorName := strings.TrimSuffix(newPath[idx+len("textures/"):], ext)
		if newName, ok := HorseArmorRenames[armorName]; ok {
			return prefix + newName + ext, true
		}
	}

	if idx := strings.Index(newPath, "textures/misc/"); idx >= 0 {
		prefix := newPath[:idx+len("textures/")]
		miscName := strings.TrimSuffix(newPath[idx+len("textures/"):], ext)
		if newName, ok := MiscTextureRenames[miscName]; ok {
			return prefix + newName + ext, true
		}
	}

	return newPath, renamed
}

func isBlockTextureDir(dir string) bool {
	return strings.HasSuffix(dir, "/textures/block") ||
		strings.HasSuffix(dir, "/textures/blocks")
}

func isItemTextureDir(dir string) bool {
	return strings.HasSuffix(dir, "/textures/item") ||
		strings.HasSuffix(dir, "/textures/items")
}

// RemapTextureReference converts a texture reference inside model JSON
// from 1.8.9 to 1.21 format (e.g. "blocks/stone" → "minecraft:block/stone").
func RemapTextureReference(ref string) string {
	cleanRef := strings.TrimPrefix(ref, "minecraft:")

	if strings.HasPrefix(cleanRef, "blocks/") {
		texName := cleanRef[len("blocks/"):]
		if newName, ok := BlockTextureRenames[texName]; ok {
			return "minecraft:block/" + newName
		}
		return "minecraft:block/" + texName
	}

	if strings.HasPrefix(cleanRef, "items/") {
		texName := cleanRef[len("items/"):]
		if newName, ok := ItemTextureRenames[texName]; ok {
			return "minecraft:item/" + newName
		}
		return "minecraft:item/" + texName
	}

	if strings.HasPrefix(cleanRef, "block/") {
		texName := cleanRef[len("block/"):]
		if newName, ok := BlockTextureRenames[texName]; ok {
			return "minecraft:block/" + newName
		}
		return "minecraft:" + cleanRef
	}

	if strings.HasPrefix(cleanRef, "item/") {
		texName := cleanRef[len("item/"):]
		if newName, ok := ItemTextureRenames[texName]; ok {
			return "minecraft:item/" + newName
		}
		return "minecraft:" + cleanRef
	}

	if strings.HasPrefix(cleanRef, "entity/") {
		entityName := cleanRef[len("entity/"):]
		if newName, ok := EntityTextureRenames[entityName]; ok {
			return "minecraft:entity/" + newName
		}
		return "minecraft:" + cleanRef
	}

	return ref
}

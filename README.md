<<<<<<< HEAD
# MC Texture Pack Converter

Converts Minecraft Java resource packs from **1.7 / 1.8.9** to **1.21.10** (`pack_format 69`).

![GUI screenshot placeholder](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-blue)

## Features

- Renames block, item, entity, armor, and misc textures to 1.21 naming
- Renames sound events (`mob.zombie.say` → `entity.zombie.ambient`, etc.)
- Converts `.lang` files to `.json`
- Updates `pack.mcmeta` with correct `pack_format` and `supported_formats`
- Splits painting atlas into individual painting files
- Splits `widgets.png` and `icons.png` into individual GUI sprites
- Splits double chest textures into left/right halves
- Fixes panorama cubemap faces (pads to square)
- Duplicates enchant glint texture for item + entity variants
- Fixes crosshair transparency for packs that used XOR blending

## Install

Grab the latest build from the [Releases](../../releases) page — just download and run, no install needed.

## Build from source

**Requirements:**
- Go 1.21+
- C compiler (GCC/MinGW on Windows, Xcode CLI tools on macOS) — needed by Fyne for OpenGL bindings
- On Linux: `libgl1-mesa-dev xorg-dev` (or equivalent for your distro)

```bash
git clone https://github.com/youruser/mc-texture-pack-converter.git
cd mc-texture-pack-converter
go build -o mc-pack-converter .
```

On Windows:
```bash
go build -o mc-pack-converter.exe .
```

### Linux dependencies (Debian/Ubuntu)

```bash
sudo apt install gcc libgl1-mesa-dev xorg-dev
```

### Linux dependencies (Fedora)

```bash
sudo dnf install gcc mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libXxf86vm-devel
```

## Usage

1. Run the executable
2. Click **Browse** to pick your input `.zip` pack
3. Output path auto-generates as `<packname>_1.21.zip` — click **Change** to override
4. Click **Convert**

The converter reads all files upfront, processes them in parallel across all CPU cores, and streams the output zip directly to disk. Pre-compressed files (PNG, OGG, etc.) are stored without re-compression.

## Project structure

```
├── main.go                  Entry point
├── gui/
│   └── gui.go               Fyne GUI with native file dialogs
└── converter/
    ├── converter.go          Core pipeline (read → parallel process → write)
    ├── atlas.go              Atlas splitting (paintings, widgets, icons, chests)
    ├── textures.go           File path remapping
    ├── models.go             Model/blockstate JSON conversion
    ├── mappings.go           All rename tables and atlas region definitions
    ├── packmeta.go           pack.mcmeta conversion
    ├── sounds.go             sounds.json event remapping
    └── lang.go               .lang → .json conversion
```

## License

MIT
=======
# PackConverter
>>>>>>> 247f3cbb019c608cd2d6cde29adde3a75bdf2fe4

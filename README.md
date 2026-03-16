# MC Texture Pack Converter

Converts Minecraft Java resource packs from **1.7 / 1.8.9** to **1.21.10** (`pack_format 69`).

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
git clone https://github.com/ratph6/PackConverter.git
cd PackConverter
go build -o PackConverter .
```

On Windows:
```bash
go build -o PackConverter.exe .
```

## Usage

1. Run the executable
2. Click **Browse** to pick your input `.zip` pack
3. Output path auto-generates as `<packname>_1.21.zip` — click **Change** to override
4. Click **Convert**

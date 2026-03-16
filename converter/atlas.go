package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"path"
	"strings"
)

// SplitPaintingAtlas splits paintings_kristoffer_zetterstrand.png into individual files.
func SplitPaintingAtlas(data []byte) (map[string][]byte, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode painting atlas: %w", err)
	}

	bounds := img.Bounds()
	if bounds.Dx() < 256 || bounds.Dy() < 256 {
		return nil, fmt.Errorf("painting atlas too small: %dx%d", bounds.Dx(), bounds.Dy())
	}

	// Standard atlas is 256x256 with a 16-unit grid. HD packs scale proportionally.
	unit := bounds.Dx() / 16

	results := make(map[string][]byte)
	for _, region := range PaintingAtlasRegions {
		x := region.X * unit
		y := region.Y * unit
		w := region.W * unit
		h := region.H * unit

		if x+w > bounds.Dx() || y+h > bounds.Dy() {
			continue
		}

		cropped := cropImage(img, x, y, w, h)
		var buf bytes.Buffer
		if err := png.Encode(&buf, cropped); err != nil {
			continue
		}
		results[region.Name+".png"] = buf.Bytes()
	}

	// back.png from top-right area of the atlas
	backX := 12 * unit
	if backX+unit <= bounds.Dx() && unit <= bounds.Dy() {
		cropped := cropImage(img, backX, 0, unit, unit)
		var buf bytes.Buffer
		if err := png.Encode(&buf, cropped); err == nil {
			results["back.png"] = buf.Bytes()
		}
	}

	return results, nil
}

// SplitWidgetAtlas splits gui/widgets.png into individual sprites.
func SplitWidgetAtlas(data []byte) (map[string][]byte, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode widgets atlas: %w", err)
	}

	bounds := img.Bounds()
	scale := bounds.Dx() / 256
	if scale < 1 {
		scale = 1
	}

	results := make(map[string][]byte)
	for _, region := range WidgetAtlasRegions {
		x := region.X * scale
		y := region.Y * scale
		w := region.W * scale
		h := region.H * scale

		if x+w > bounds.Dx() || y+h > bounds.Dy() {
			continue
		}

		cropped := cropImage(img, x, y, w, h)
		var buf bytes.Buffer
		if err := png.Encode(&buf, cropped); err != nil {
			continue
		}
		results[region.OutputPath+".png"] = buf.Bytes()
	}

	return results, nil
}

// SplitIconsAtlas splits gui/icons.png into individual HUD sprites.
func SplitIconsAtlas(data []byte) (map[string][]byte, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode icons atlas: %w", err)
	}

	bounds := img.Bounds()
	scale := bounds.Dx() / 256
	if scale < 1 {
		scale = 1
	}

	results := make(map[string][]byte)
	for _, region := range IconsAtlasRegions {
		x := region.X * scale
		y := region.Y * scale
		w := region.W * scale
		h := region.H * scale

		if x+w > bounds.Dx() || y+h > bounds.Dy() {
			continue
		}

		cropped := cropImageNRGBA(img, x, y, w, h)

		// 1.8.9 rendered the crosshair via XOR blending, so many packs have
		// opaque black backgrounds that were invisible. 1.21+ uses alpha
		// blending, which makes those black pixels show up as squares.
		fixOpaqueBlackBackground(cropped)

		var buf bytes.Buffer
		if err := png.Encode(&buf, cropped); err != nil {
			continue
		}
		results[region.OutputPath+".png"] = buf.Bytes()
	}

	return results, nil
}

// fixOpaqueBlackBackground makes solid-black pixels transparent if the image
// appears to have a non-transparent black background (>=2 corners are opaque black).
func fixOpaqueBlackBackground(img *image.NRGBA) {
	bounds := img.Bounds()
	corners := []image.Point{
		{bounds.Min.X, bounds.Min.Y},
		{bounds.Max.X - 1, bounds.Min.Y},
		{bounds.Min.X, bounds.Max.Y - 1},
		{bounds.Max.X - 1, bounds.Max.Y - 1},
	}
	blackCorners := 0
	for _, pt := range corners {
		c := img.NRGBAAt(pt.X, pt.Y)
		if c.R == 0 && c.G == 0 && c.B == 0 && c.A == 255 {
			blackCorners++
		}
	}
	if blackCorners < 2 {
		return
	}

	transparent := color.NRGBA{0, 0, 0, 0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBAAt(x, y)
			if c.R == 0 && c.G == 0 && c.B == 0 && c.A == 255 {
				img.SetNRGBA(x, y, transparent)
			}
		}
	}
}

// SplitDoubleChest splits a 1.8.9 double chest texture into left/right halves.
func SplitDoubleChest(data []byte) (left []byte, right []byte, err error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode chest texture: %w", err)
	}

	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	halfW := w / 2

	leftImg := cropImage(img, 0, 0, halfW, h)
	rightImg := cropImage(img, halfW, 0, halfW, h)

	var leftBuf, rightBuf bytes.Buffer
	if err := png.Encode(&leftBuf, leftImg); err != nil {
		return nil, nil, err
	}
	if err := png.Encode(&rightBuf, rightImg); err != nil {
		return nil, nil, err
	}

	return leftBuf.Bytes(), rightBuf.Bytes(), nil
}

func isPanoramaFace(p string) bool {
	if !strings.Contains(p, "textures/gui/title/background/") {
		return false
	}
	base := path.Base(p)
	for i := 0; i <= 5; i++ {
		if base == fmt.Sprintf("panorama_%d.png", i) {
			return true
		}
	}
	return false
}

// squarePanoramaFace pads a panorama face to be square. Returns nil if already square.
func squarePanoramaFace(data []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode panorama face: %w", err)
	}

	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if w == h {
		return nil, nil
	}

	size := w
	if h > w {
		size = h
	}

	dst := image.NewNRGBA(image.Rect(0, 0, size, size))
	offsetX := (size - w) / 2
	offsetY := (size - h) / 2
	draw.Draw(dst, image.Rect(offsetX, offsetY, offsetX+w, offsetY+h), img, bounds.Min, draw.Src)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func cropImage(src image.Image, x, y, w, h int) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(dst, dst.Bounds(), src, image.Pt(x, y), draw.Src)
	return dst
}

func cropImageNRGBA(src image.Image, x, y, w, h int) *image.NRGBA {
	dst := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(dst, dst.Bounds(), src, image.Pt(x, y), draw.Src)
	return dst
}

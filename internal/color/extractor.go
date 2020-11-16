package color

import (
	"fmt"
	"github.com/gondo/pex-interview/internal/color_extractor"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ExtractColors(img image.Image, count int) (colors []string) {
	if img == nil {
		return colors
	}

	extractedColors := color_extractor.ExtractColors(img, count)

	for _, c := range extractedColors {
		r, g, b, _ := c.RGBA()
		colorString := fmt.Sprintf("#%s%s%s", colorToHex(r), colorToHex(g), colorToHex(b))
		colors = append(colors, colorString)
	}

	return colors
}

func colorToHex(color uint32) string {
	return fmt.Sprintf("%02X", color/0x101)
}

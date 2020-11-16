package color

import (
	"image"
	"image/color"
	"testing"
)

func TestDetectColors(t *testing.T) {
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	testCases := map[string]struct {
		img    image.Image
		count  int
		colors []string
	}{
		"Empty image": {
			img:    nil,
			count:  3,
			colors: []string{},
		},
		"1 color image": {
			img: imageFromColors([]color.Color{
				red,
			}),
			count: 3,
			colors: []string{
				"#FF0000",
			},
		},
		"2 colors image": {
			img: imageFromColors([]color.Color{
				red,
				green,
			}),
			count: 3,
			colors: []string{
				"#00FF00",
				"#FF0000",
			},
		},
		"3 colors image, 1 extracted": {
			img: imageFromColors([]color.Color{
				red,
				green,
				blue,
			}),
			count: 1,
			colors: []string{
				"#0000FF",
			},
		},
		"3 colors image, green extracted": {
			img: imageFromColors([]color.Color{
				red,
				green,
				blue,
				green,
			}),
			count: 1,
			colors: []string{
				"#00FF00",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			colors := ExtractColors(testCase.img, testCase.count)
			if !equal(testCase.colors, colors) {
				t.Errorf("%v expected, got %v", testCase.colors, colors)
			}
		})
	}
}

func TestColorToHex(t *testing.T) {
	testCases := []struct {
		color   uint32
		expected string
	}{
		{
			color: 0,
			expected: "00",
		},
		{
			color: 65535,
			expected: "FF",
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			hex := colorToHex(testCase.color)
			if testCase.expected != hex {
				t.Errorf("%v expected, got %v", testCase.expected, hex)
			}
		})
	}
}

func imageFromColors(colors []color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, len(colors), 1))
	for i, c := range colors {
		img.Set(i, 0, c)
	}
	return img
}

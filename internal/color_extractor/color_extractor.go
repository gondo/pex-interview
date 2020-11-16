package color_extractor

import (
	"image"
	"image/color"
	"math"
	"sort"
)

type bucket struct {
	Red   float64
	Green float64
	Blue  float64
	Count float64
}

type byCount []bucket

func (c byCount) Len() int           { return len(c) }
func (c byCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCount) Less(i, j int) bool { return c[i].Count < c[j].Count }

func ExtractColors(image image.Image, count int) []color.Color {
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	stepX := 1
	stepY := 1

	// load image's pixels into buckets
	var buckets [2][2][2]bucket
	totalCount := 0.
	for x := 0; x < width; x += stepX {
		for y := 0; y < height; y += stepY {
			c := image.At(x, y)
			r, g, b, a := c.RGBA()
			r >>= 8
			g >>= 8
			b >>= 8
			a >>= 8
			i := r >> 7
			j := g >> 7
			k := b >> 7
			if a > 0 {
				alphaFactor := float64(a) / 255.
				buckets[i][j][k].Red += float64(r) * alphaFactor
				buckets[i][j][k].Green += float64(g) * alphaFactor
				buckets[i][j][k].Blue += float64(b) * alphaFactor
				buckets[i][j][k].Count += alphaFactor
				totalCount += alphaFactor
			}
		}
	}

	// calculate bucket's averages
	var bucketsAverages []bucket
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				currentBucket := buckets[i][j][k]
				if currentBucket.Count > 0 {
					bucketsAverages = append(bucketsAverages, bucket{
						Count: currentBucket.Count,
						Red:   currentBucket.Red / currentBucket.Count,
						Green: currentBucket.Green / currentBucket.Count,
						Blue:  currentBucket.Blue / currentBucket.Count,
					})
				}
			}
		}
	}

	// sort buckets by bucket size
	sort.Sort(sort.Reverse(byCount(bucketsAverages)))

	// export color.Color from bucket
	colors := []color.Color{}
	for _, avg := range bucketsAverages {
		if len(colors) >= count {
			break
		}
		colors = append(colors, color.RGBA{
			R: uint8(math.Round(avg.Red)),
			G: uint8(math.Round(avg.Green)),
			B: uint8(math.Round(avg.Blue)),
			A: 255,
		})
	}

	return colors
}

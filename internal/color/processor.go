package color

import (
	"fmt"
	"github.com/gondo/pex-interview/internal"
	"github.com/gondo/pex-interview/internal/logger"
	"strings"
)

// NEW
func Process2(images <-chan internal.DownloadedImage, results chan<- string, count, id int) {
	for dImg := range images {
		logger.Log(fmt.Sprintf("extracting colors from %s", dImg.Url))

		colors := ExtractColors(dImg.Img, count)
		colors = fillMissing(colors, count)
		results <- OutputFormat(dImg.Url, colors)
	}

	logger.Log(fmt.Sprintf("extracting colors %d done", id))
}

// Extract colors from an image and push result to a channel
func Process(dImg internal.DownloadedImage, results chan<- string, colorsCount int) {
	logger.Log(fmt.Sprintf("extracting colors from %s", dImg.Url))

	colors := ExtractColors(dImg.Img, colorsCount)
	colors = fillMissing(colors, colorsCount)
	results <- OutputFormat(dImg.Url, colors)
}

func OutputFormat(url string, colors []string) string {
	return fmt.Sprintf("%s,%s", url, strings.Join(colors, ","))
}

func fillMissing(colors []string, count int) []string {
	for i := len(colors); i < count; i++ {
		colors = append(colors, "")
	}
	return colors
}

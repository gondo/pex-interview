package downloader

import (
	"fmt"
	"github.com/gondo/pex-interview/internal"
	"github.com/gondo/pex-interview/internal/color"
	"github.com/gondo/pex-interview/internal/logger"
	"sync"
)

func Process(urls <-chan string, results chan<- string, colorsCount, id int) {
	// Starts additional WaitGroup for workers extracting colors from downloaded images.
	var wg sync.WaitGroup

	for url := range urls {
		logger.Log(fmt.Sprintf("downloader %d url %s", id, url))

		img, err := DownloadImage(url)
		if nil != err {
			logger.Log(fmt.Sprintf("downloader %d error %s", id, err))
		}

		// Bypass color detection
		if nil == img {
			results <- color.OutputFormat(url, []string{})
			continue
		}

		dImg := internal.DownloadedImage{
			Img: img,
			Url: url,
		}

		wg.Add(1)
		go func(dImg internal.DownloadedImage, results chan<- string, colorsCount int) {
			defer wg.Done()
			color.Process(dImg, results, colorsCount)
		}(dImg, results, colorsCount)
	}

	wg.Wait()
	logger.Log(fmt.Sprintf("dowloader %d done", id))
}

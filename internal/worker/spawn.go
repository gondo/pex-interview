package worker

import (
	"github.com/gondo/pex-interview/internal/downloader"
	"sync"
)

func SpawnWorkers(count int, wg *sync.WaitGroup, colorsCount int) (chan string, chan string) {
	// Will read only a limited number of urls from input, one for each worker.
	urls := make(chan string, count)

	// Extracted colors will be pushed into a `results` channel from which they will be stored into an output.
	// Storing results is expected to take less time then downloading and extracting combined,
	// Therefore channel size is set to `count`, which is more than enough.
	results := make(chan string, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(urls <-chan string, results chan<- string, colorsCount, i int) {
			defer wg.Done()
			downloader.Process(urls, results, colorsCount, i)
		}(urls, results, colorsCount, i)
	}

	return urls, results
}
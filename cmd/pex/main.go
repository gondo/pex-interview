package main

import (
	"fmt"
	"github.com/gondo/pex-interview/internal"
	"github.com/gondo/pex-interview/internal/input"
	"github.com/gondo/pex-interview/internal/logger"
	"github.com/gondo/pex-interview/internal/output"
	"github.com/gondo/pex-interview/internal/worker"
	"sync"
	"time"
)

// Number of colors to be extracted
func main() {
	p := internal.ParseFlags()

	if p.Debug {
		logger.EnableDebug()
		logger.Log(fmt.Sprintf("Dowloading %s into %s using %d workers\n", p.Input, p.Output, p.NumberOfWorkers))

		start := time.Now()
		defer func() {
			elapsed := time.Since(start)
			logger.Log(fmt.Sprintf("duration: %s \n", elapsed))
		}()
	}

	run(p)
}

func run(p internal.Params) {
	var wg sync.WaitGroup
	urls, results := worker.SpawnWorkers(p.NumberOfWorkers, &wg, p.NumberOfColors)

	go input.ReadFileToChannel(p.Input, urls)
	go wait(&wg, results)

	output.StoreToFile(p.Output, results)
}

// Waiting for workers to finish so we can close `results` channel
func wait(wg *sync.WaitGroup, results chan<- string) {
	wg.Wait()
	close(results)
}

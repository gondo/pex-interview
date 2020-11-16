package input

import (
	"bufio"
	"github.com/gondo/pex-interview/internal"
	"github.com/gondo/pex-interview/internal/logger"
	"os"
)

// Input file is read line by line and each line is pushed into a channel.
func ReadFileToChannel(fileName string, urls chan<- string) {
	file, err := os.Open(fileName)
	internal.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		urls <- url
	}

	if err := scanner.Err(); err != nil {
		internal.CheckError(err)
	}

	close(urls)
	logger.Log("reading input done")
}

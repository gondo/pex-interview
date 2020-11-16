package output

import (
	"github.com/gondo/pex-interview/internal"
	"github.com/gondo/pex-interview/internal/logger"
	"os"
)

// Read results from channel and store to file
func StoreToFile(fileName string, results chan string) {
	file, err := os.Create(fileName)
	internal.CheckError(err)
	defer file.Close()

	first := true
	for res := range results {
		// Prevents extra newline at the end of a file
		if first {
			first = false
		} else {
			res = "\n" + res
		}
		_, err = file.WriteString(res)
		internal.CheckError(err)
	}

	logger.Log("storing done")
}
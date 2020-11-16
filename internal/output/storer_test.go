package output

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestStore(t *testing.T) {
	ch := make(chan string, 10)
	input := []string{"a", "b", "c"}
	for _, s := range input {
		ch <- s
	}
	close(ch)
	fileName := "test_data/test_output.csv"
	StoreToFile(fileName, ch)

	output := readFile(fileName)
	if !equal(input, output) {
		t.Errorf("%v expected, got %v", input, output)
	}
}

func readFile(fileName string) []string {
	b, _ := ioutil.ReadFile(fileName)
	return strings.Split(string(b), "\n")
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
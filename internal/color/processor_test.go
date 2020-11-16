package color

import (
	"github.com/gondo/pex-interview/internal"
	"testing"
)

func TestOutputFormat(t *testing.T) {
	out := OutputFormat("http://www.pex.com", []string{"#FF0000", "#00FF00", "#0000FF"})
	expected := "http://www.pex.com,#FF0000,#00FF00,#0000FF"
	if out != expected {
		t.Errorf("%v expected, got %v", expected, out)
	}
}

func TestFillMissing(t *testing.T) {
	testCases := []struct {
		colors   []string
		count    int
		expected []string
	}{
		{
			colors:   []string{},
			count:    3,
			expected: []string{"", "", ""},
		},
		{
			colors:   []string{"#FF0000", "#00FF00", "#0000FF"},
			count:    1,
			expected: []string{"#FF0000", "#00FF00", "#0000FF"},
		},
		{
			colors:   []string{"#FF0000"},
			count:    3,
			expected: []string{"#FF0000", "", ""},
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			out := fillMissing(testCase.colors, testCase.count)
			if !equal(testCase.expected, out) {
				t.Errorf("%v expected, got %v", testCase.expected, out)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	ch := make(chan string, 2)
	Process(internal.DownloadedImage{}, ch, 1)
	if len(ch) != 1 {
		t.Fatal("channel doesn't have expected items")
	}
}
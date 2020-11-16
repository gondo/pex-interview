package input

import (
	"testing"
)

func TestReadFileToChannel(t *testing.T) {
	ch := make(chan string, 10)
	ReadFileToChannel("test_data/test_input.txt", ch)

	var out []string
	for s := range ch {
		out = append(out, s)
	}

	expected := []string{
		"http://i.imgur.com/FApqk3D.jpg",
		"http://i.imgur.com/TKLs9lo.jpg",
		"https://i.redd.it/d8021b5i2moy.jpg",
		"https://i.redd.it/4m5yk8gjrtzy.jpg",
		"https://i.redd.it/xae65ypfqycy.jpg",
	}
	if !equal(out, expected) {
		t.Fatal("channel doesn't have expected items")
	}
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

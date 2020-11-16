package internal

import (
	"flag"
	"fmt"
	"os"
)

type Params struct {
	Debug           bool
	Input           string
	NumberOfWorkers int
	NumberOfColors  int
	Output          string
}

var p Params

func ParseFlags() Params {
	flag.IntVar(&p.NumberOfColors, "c", 3, "Number of colors to extract")
	flag.BoolVar(&p.Debug, "d", false, "Debug output")
	flag.IntVar(&p.NumberOfWorkers, "n", 100, "Number of workers")

	flag.Usage = func() {
		fmt.Println("Tool for extracting the most prevalent colors from a given list of image URLs.")
		fmt.Println("by Michal Gondar 2020")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("	pex [flags] input-path output-path \n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	p.Input = flag.Arg(0)
	p.Output = flag.Arg(1)

	return p
}

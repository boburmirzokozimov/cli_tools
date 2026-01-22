package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	opts := DisplayOptions{}
	flag.BoolVar(&opts.Words, "w", false, "print word count")
	flag.BoolVar(&opts.Lines, "l", false, "print line count")
	flag.BoolVar(&opts.Bytes, "c", false, "print byte count")
	flag.Parse()

	total := Counts{}
	fileNames := flag.Args()
	didErr := false
	for _, fileName := range fileNames {
		counts, err := CountFile(fileName)
		if err != nil {
			fmt.Println("counter: ", err)
			didErr = true
			continue
		}

		counts.PrintWithOptions(os.Stdout, opts, fileName)
		total.Add(&counts)
	}
	if len(fileNames) > 1 {
		total.PrintWithOptions(os.Stdout, opts, "total")
	}

	if len(fileNames) == 0 {
		GetCounts(os.Stdin).PrintWithOptions(os.Stdout, opts)
	}
	if didErr {
		os.Exit(1)
	}

}

package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	opts := DisplayOptions{}
	flag.BoolVar(&opts.Words, "w", false, "print word count")
	flag.BoolVar(&opts.Lines, "l", false, "print line count")
	flag.BoolVar(&opts.Bytes, "c", false, "print byte count")
	flag.Parse()

	total := Counts{}
	fileNames := flag.Args()
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	didErr := false
	for _, fileName := range fileNames {
		counts, err := CountFile(fileName)
		if err != nil {
			fmt.Println("counter: ", err)
			didErr = true
			continue
		}

		counts.PrintWithOptions(writer, opts, fileName)
		total.Add(&counts)
	}
	if len(fileNames) > 1 {
		total.PrintWithOptions(writer, opts, "total")
	}

	if len(fileNames) == 0 {
		GetCounts(os.Stdin).PrintWithOptions(writer, opts)
	}
	if didErr {
		writer.Flush()
		os.Exit(1)
	}
	writer.Flush()

}

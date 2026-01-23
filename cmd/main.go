package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	counter "github.com/boburmirzokozimov/cli_tools"
	"github.com/boburmirzokozimov/cli_tools/display"
)

func main() {
	args := display.NewOptionArgs()

	opts := display.NewOptions(args)

	flag.BoolVar(&args.Words, "w", false, "print word count")
	flag.BoolVar(&args.Lines, "l", false, "print line count")
	flag.BoolVar(&args.Bytes, "c", false, "print byte count")
	flag.Parse()

	total := counter.Counts{}
	fileNames := flag.Args()
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	didErr := false
	for _, fileName := range fileNames {
		counts, err := counter.CountFile(fileName)
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
		counter.GetCounts(os.Stdin).PrintWithOptions(writer, opts)
	}
	if didErr {
		writer.Flush()
		os.Exit(1)
	}
	writer.Flush()

}

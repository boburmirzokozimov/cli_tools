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

	flag.BoolVar(&args.Words, "w", false, "print word count")
	flag.BoolVar(&args.Lines, "l", false, "print line count")
	flag.BoolVar(&args.Bytes, "c", false, "print byte count")
	flag.Parse()

	opts := display.NewOptions(args).WithDefaults()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	didErr := false
	fileNames := flag.Args()
	total := counter.Counts{}

	ch := counter.CountFiles(fileNames)

	for res := range ch {
		total.Add(&res.Counts)
		res.Counts.PrintWithOptions(writer, opts, res.Filename)
		if res.Err != nil {
			fmt.Println("counter: ", res.Err)
			didErr = true
		}
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

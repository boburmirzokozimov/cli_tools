package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
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

	wg := sync.WaitGroup{}
	wg.Add(len(fileNames))
	l := sync.Mutex{}

	for _, fileName := range fileNames {
		go func(name string) {
			defer wg.Done()

			counts, err := counter.CountFile(name)
			if err != nil {
				l.Lock()
				fmt.Println("counter: ", err)
				didErr = true
				l.Unlock()
				return
			}
			l.Lock()
			defer l.Unlock()
			total.Add(&counts)
			counts.PrintWithOptions(writer, opts, name)
		}(fileName)
	}
	wg.Wait()
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

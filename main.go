package main

import (
	"fmt"
	"os"
)

func main() {

	total := Counts{}
	fileNames := os.Args[1:]
	didErr := false
	for _, fileName := range fileNames {
		counts, err := CountFile(fileName)
		if err != nil {
			fmt.Println("counter: ", err)
			didErr = true
			continue
		}

		fmt.Println(counts, fileName)
		total.words += counts.words
		total.lines += counts.lines
		total.bytes += counts.bytes
	}
	if len(fileNames) > 1 {
		PrintCounts(Counts{total.words, total.lines, total.bytes}, "total")
		fmt.Println(total, "total")
	}

	if len(fileNames) == 0 {
		counts := GetCounts(os.Stdin)

		PrintCounts(Counts{counts.words, counts.lines, counts.bytes}, "stdin")
	}
	if didErr {
		os.Exit(1)
	}

}

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

		counts.Print(os.Stdout, fileName)
		total.Add(&counts)
	}
	if len(fileNames) > 1 {
		total.Print(os.Stdout, "total")
	}

	if len(fileNames) == 0 {
		GetCounts(os.Stdin).Print(os.Stdout)
	}
	if didErr {
		os.Exit(1)
	}

}

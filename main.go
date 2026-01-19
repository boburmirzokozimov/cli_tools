package main

import (
	"fmt"
	"os"
)

func main() {

	total := 0
	fileNames := os.Args[1:]
	didErr := false
	for _, fileName := range fileNames {
		wordsCount, err := CountWordsInFile(fileName)
		if err != nil {
			fmt.Println("counter: ", err)
			didErr = true
			continue
		}

		fmt.Println(wordsCount, fileName)
		total += wordsCount
	}
	if len(fileNames) > 1 {
		fmt.Println(total, "total")
	}

	if len(fileNames) == 0 {
		wordsCount := CountWords(os.Stdin)
		fmt.Println(wordsCount, "stdin")
	}
	if didErr {
		os.Exit(1)
	}

}

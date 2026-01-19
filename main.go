package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Should provide at least one filename")
	}
	total := 0
	fileNames := os.Args[1:]
	didErr := false
	for _, fileName := range fileNames {
		wordsCount, err := CountWordsInFile(fileName)
		if err != nil {
			fmt.Println(os.Stderr, "counter: ", err)
			didErr = true
			continue
		}

		fmt.Println(wordsCount, fileName)
		total += wordsCount
	}
	if len(fileNames) > 1 {
		fmt.Println(total, "total")
	}
	if didErr {
		os.Exit(1)
	}

}

func CountWordsInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("open file %w", err)
	}
	defer file.Close()
	return CountWords(file), nil
}

func CountWords(file io.Reader) int {
	wordsCount := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordsCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("failed to scan file:", err)
	}
	return wordsCount
}

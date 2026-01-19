package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

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

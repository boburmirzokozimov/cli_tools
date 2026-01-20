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

func CountLines(file io.Reader) int {
	linesCount := 0
	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r == '\n' {
			linesCount++
		}
	}

	return linesCount
}

func CountBytes(file io.Reader) int {
	cnt, _ := io.Copy(io.Discard, file)
	return int(cnt)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Counts struct {
	words int
	lines int
	bytes int
}

func GetCounts(file io.ReadSeeker) Counts {
	count := Counts{}
	const offSet = 0

	count.words = CountWords(file)
	file.Seek(offSet, io.SeekStart)

	count.lines = CountLines(file)
	file.Seek(offSet, io.SeekStart)

	count.bytes = CountBytes(file)

	return count
}

func CountFile(fileName string) (Counts, error) {
	file, err := os.Open(fileName)
	count := Counts{}
	if err != nil {
		return count, fmt.Errorf("open file %w", err)
	}

	defer file.Close()

	count = GetCounts(file)

	return count, nil
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

func PrintCounts(counts Counts, filename string) {
	fmt.Println(counts.words, counts.lines, counts.bytes, filename)
}

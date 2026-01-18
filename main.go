package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal("failed to read file:", err)
	}
	defer file.Close()
	wordsCount := CountWordsInFile(file)
	fmt.Println(wordsCount)
}

func CountWords(data []byte) int {
	return len(bytes.Fields(data))
}

func CountWordsInFile(file *os.File) int {
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

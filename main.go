package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
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
	isInsideWord := false
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if !unicode.IsSpace(r) && !isInsideWord {
			wordsCount++
		}

		isInsideWord = !unicode.IsSpace(r)
	}

	return wordsCount
}

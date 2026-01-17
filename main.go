package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("./words.txt")
	if err != nil {
		return
	}

	counter := countWords(file)
	fmt.Println(counter)
}

func countWords(data []byte) int {
	counter := 1
	for i := range data {
		if data[i] == ' ' {
			counter++
		}

	}

	return counter
}

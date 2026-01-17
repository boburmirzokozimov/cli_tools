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
	counter := 0
	if len(data) > 0 {
		counter++
	}
	for i := range data {
		if data[i] == ' ' {
			counter++
		}

	}

	return counter
}

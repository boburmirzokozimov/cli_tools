package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("./words.txt")
	if err != nil {
		return
	}

	counter := CountWords(file)
	fmt.Println(counter)
}

func CountWords(data []byte) int {
	return len(bytes.Fields(data))
}

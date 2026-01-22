package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Counts struct {
	words int
	lines int
	bytes int
}

type DisplayOptions struct {
	Words bool
	Lines bool
	Bytes bool
}

func (opts DisplayOptions) withDefaults() DisplayOptions {
	if !opts.Words && !opts.Lines && !opts.Bytes {
		return DisplayOptions{Words: true, Lines: true, Bytes: true}
	}
	return opts
}

func GetCounts(file io.Reader) Counts {
	count := Counts{}
	reader := bufio.NewReader(file)
	isInsideWord := false
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			count.lines++
		}
		count.bytes += size

		isSpace := unicode.IsSpace(r)

		if !isInsideWord && !isSpace {
			count.words++
		}

		isInsideWord = !isSpace
	}
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

func (this Counts) Print(w io.Writer, filename ...string) {
	this.PrintWithOptions(w, DisplayOptions{Words: true, Lines: true, Bytes: true}, filename...)
}

func (this Counts) PrintWithOptions(w io.Writer, opts DisplayOptions, filename ...string) {
	opts = opts.withDefaults()
	fields := make([]string, 0, 3)
	if opts.Words {
		fields = append(fields, strconv.Itoa(this.words))
	}
	if opts.Lines {
		fields = append(fields, strconv.Itoa(this.lines))
	}
	if opts.Bytes {
		fields = append(fields, strconv.Itoa(this.bytes))
	}

	fmt.Fprint(w, strings.Join(fields, " "))

	for _, name := range filename {
		fmt.Fprintf(w, " %s", name)
	}

	fmt.Fprintf(w, "\n")
}

func (this *Counts) Add(other *Counts) {
	this.words += other.words
	this.lines += other.lines
	this.bytes += other.bytes
}

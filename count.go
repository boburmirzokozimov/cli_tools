package counter

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/boburmirzokozimov/cli_tools/display"
)

type Counts struct {
	words int
	lines int
	bytes int
}

type FileCountsResult struct {
	Counts   Counts
	Filename string
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

func (this Counts) PrintWithOptions(w io.Writer, opts display.DisplayOptions, filename ...string) {

	fields := make([]string, 0, 3)
	if opts.ShowWords() {
		fields = append(fields, strconv.Itoa(this.words))
	}
	if opts.ShowLines() {
		fields = append(fields, strconv.Itoa(this.lines))
	}
	if opts.ShowBytes() {
		fields = append(fields, strconv.Itoa(this.bytes))
	}

	fmt.Fprint(w, strings.Join(fields, "\t"))

	for _, name := range filename {
		fmt.Fprintf(w, "\t%s", name)
	}

	fmt.Fprintf(w, "\n")
}

func (this *Counts) Add(other *Counts) {
	this.words += other.words
	this.lines += other.lines
	this.bytes += other.bytes
}

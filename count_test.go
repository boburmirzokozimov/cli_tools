package main

import (
	"bytes"
	"testing"
)

func TestGetCounts(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect Counts
	}{
		{name: "empty input", input: "", expect: Counts{words: 0, lines: 0, bytes: 0}},
		{name: "single space", input: " ", expect: Counts{words: 0, lines: 0, bytes: 1}},
		{name: "multiple words", input: "one two three four five", expect: Counts{words: 5, lines: 0, bytes: 23}},
		{name: "leading and trailing spaces", input: "  one two  ", expect: Counts{words: 2, lines: 0, bytes: 11}},
		{name: "mixed separators", input: "one\ntwo\tthree", expect: Counts{words: 3, lines: 1, bytes: 13}},
		{name: "multiple lines with trailing newline", input: "one\ntwo\nthree\n", expect: Counts{words: 3, lines: 3, bytes: 14}},
		{name: "windows line endings", input: "one\r\ntwo\r\nthree\r\n", expect: Counts{words: 3, lines: 3, bytes: 17}},
		{name: "line with spaces", input: "   \n", expect: Counts{words: 0, lines: 1, bytes: 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tc.input))
			result := GetCounts(reader)
			if result != tc.expect {
				t.Fatalf("expected %+v, got %+v", tc.expect, result)
			}
		})
	}
}

func TestCountsAdd(t *testing.T) {
	left := Counts{words: 2, lines: 1, bytes: 5}
	right := Counts{words: 3, lines: 4, bytes: 7}

	left.Add(&right)

	expect := Counts{words: 5, lines: 5, bytes: 12}
	if left != expect {
		t.Fatalf("expected %+v, got %+v", expect, left)
	}
}

func TestCountsPrintWithOptions(t *testing.T) {
	counts := Counts{words: 2, lines: 3, bytes: 5}
	tests := []struct {
		name    string
		options DisplayOptions
		expect  string
	}{
		{name: "defaults when none selected", options: DisplayOptions{}, expect: "2\t3\t5\n"},
		{name: "words only", options: DisplayOptions{Words: true}, expect: "2\n"},
		{name: "lines and bytes", options: DisplayOptions{Lines: true, Bytes: true}, expect: "3\t5\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			counts.PrintWithOptions(&buf, tc.options)
			if buf.String() != tc.expect {
				t.Fatalf("expected %q, got %q", tc.expect, buf.String())
			}
		})
	}
}

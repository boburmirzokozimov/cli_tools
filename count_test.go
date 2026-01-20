package main_test

import (
	"strings"
	"testing"

	counter "github.com/boburmirzokozimov/cli_tools"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{name: "multiple words", input: "one two three four five", expect: 5},
		{name: "empty string", input: "", expect: 0},
		{name: "single space", input: " ", expect: 0},
		{name: "new lines", input: "one two three", expect: 3},
		{name: "leading and trailing spaces", input: "  one two  ", expect: 2},
		{name: "multiple spaces between words", input: "one   two   three", expect: 3},
		{name: "tab separated", input: "one	two	three", expect: 3},
		{name: "mixed separators", input: "one\ntwo\tthree", expect: 3},
		{name: "mixed separators with spaces", input: "one\ntwo\tthree", expect: 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountWords(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}
func TestCountLines(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{name: "empty input", input: "", expect: 0},
		{name: "single line without newline", input: "one two three", expect: 0},
		{name: "single line with newline", input: "one two three\n", expect: 1},
		{name: "multiple lines", input: "one\ntwo\nthree", expect: 2},
		{name: "multiple lines with trailing newline", input: "one\ntwo\nthree\n", expect: 3},
		{name: "windows line endings", input: "one\r\ntwo\r\nthree\r\n", expect: 3},
		{name: "line with spaces", input: "   \n", expect: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountLines(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{name: "empty input", input: "", expect: 0},
		{name: "ascii text", input: "hello", expect: 5},
		{name: "with newline", input: `hello
`, expect: 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountBytes(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

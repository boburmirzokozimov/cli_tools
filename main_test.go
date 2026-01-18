package main_test

import (
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
			result := counter.CountWords([]byte(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

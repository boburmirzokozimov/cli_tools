package main

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{name: "multiple words", input: "one two three four five", expect: 5},
		{name: "empty string", input: "", expect: 0},
		{name: "single space", input: " ", expect: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := countWords([]byte(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

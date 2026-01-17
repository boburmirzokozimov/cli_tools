package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	expect := 5

	result := countWords([]byte(input))

	if result != expect {
		t.Fail()
	}

}

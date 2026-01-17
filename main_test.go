package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	expect := 5

	result := countWords([]byte(input))

	if result != expect {
		t.Logf("expected %d, got %d", expect, result)
		t.Fail()
	}

	input = ""
	expect = 0

	result = countWords([]byte(input))

	if result != expect {
		t.Logf("expected %d, got %d", expect, result)
		t.Fail()
	}

	input = " "
	expect = 1

	result = countWords([]byte(input))

	if result != expect {
		t.Logf("expected %d, got %d", expect, result)
		t.Fail()
	}
}

package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "This is a GREAT Time",
			expected: []string{"this", "is", "a", "great", "time"},
		},
		{
			input:    "  hello    world    ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length does not match %v actual, expected: %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("words do not match actual word %v vs expected word %v", word, expected_word)
			}
		}
	}
}

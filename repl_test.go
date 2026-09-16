package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  PIKACHU sharizard",
			expected: []string{"pikachu", "sharizard"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected is not the right len")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word %v not equal to actual word %v", expectedWord, word)
			}
		}
	}
}

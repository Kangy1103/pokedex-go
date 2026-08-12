package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello    world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "charManDer",
			expected: []string{"charmander"},
		},
		{
			input:    "I LOVE POKEMON",
			expected: []string{"i", "love", "pokemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v, expected %v", c.input, actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q) returned %v, expected %v", c.input, actual, c.expected)
				return
			}
		}
	}
}

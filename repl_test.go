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
	}

	for _, c := cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v, expected %v", c.input, actual, c.expected)
			continue
		}
		for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
			if word != expectedWord {
				t.errorf("cleanInput(%q) returned %v, expected %v", c.input, actual, c.expected)
			return "Tests failed"
			}
		}
	}
}

package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"normal spacing": {
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		"extra outside spacing": {
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		"extra middle space": {
			input:    "hello		world",
			expected: []string{"hello", "world"},
		},
		"extra middle and outside space": {
			input:    "    hello		world    ",
			expected: []string{"hello", "world"},
		},
		"capitalization": {
			input:    "Hello wOrlD",
			expected: []string{"hello", "world"},
		},
	}

	for name, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\nname: %v\n\texpected: %#v\n\tgot: %#v\n\n", name, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("\nname: %v\n\texpected: %#v\n\tgot: %#v\n\n", name, expectedWord, word)
			}
		}
	}
}

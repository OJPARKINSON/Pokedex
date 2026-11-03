package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "super,bonk test",
			expected: []string{"super,bonk", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fail()
		}

		if !slices.Equal(actual, c.expected) {
			t.Errorf("actual %s does not match expected %s", actual, c.expected)
		}
	}
}

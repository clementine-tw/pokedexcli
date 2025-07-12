package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "one word",
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			name:     "single space",
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple space",
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "have upper case",
			input:    "   Hello   worlD   ",
			expected: []string{"hello", "world"},
		},
	}

	compareSlices := func(t testing.TB, actual, expected []string) {
		t.Helper()

		if len(actual) != len(expected) {
			t.Errorf("lengths don't match: %q vs %q", actual, expected)
			return
		}
		for i, word := range actual {
			expectedWord := expected[i]
			if word != expected[i] {
				t.Errorf("expected %q, but actual is %q", expectedWord, word)
			}
		}
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Run(c.name, func(t *testing.T) {
			compareSlices(t, actual, c.expected)
		})
	}
}

package convert

import (
	"testing"
)

func TestIsStringCapitalized(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello", true},
		{"hello", false},
		{"", false},            // Empty string case
		{"H", true},            // Single capital letter
		{"h", false},           // Single lowercase letter
		{"123abc", false},      // Starts with a digit
		{"Hello world", true},  // Capitalized first word
		{"hello World", false}, // Lowercase first word
		{"!Hello", false},      // Starts with a punctuation
	}

	for _, test := range tests {
		result := IsStringCapitalized(test.input)
		if result != test.expected {
			t.Errorf("IsStringCapitalized(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

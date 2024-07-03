package convert

import (
	"testing"
)

func TestStringToUint64(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    uint64
		shouldPanic bool
	}{
		{
			name:        "Valid small number",
			input:       "123",
			expected:    123,
			shouldPanic: false,
		},
		{
			name:        "Valid large number",
			input:       "18446744073709551615", // Max uint64
			expected:    18446744073709551615,
			shouldPanic: false,
		},
		{
			name:        "Number exceeding uint64 range",
			input:       "18446744073709551616", // Max uint64 + 1
			expected:    0,
			shouldPanic: true,
		},
		{
			name:        "Invalid number format",
			input:       "invalid",
			expected:    0,
			shouldPanic: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    0,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.shouldPanic {
						t.Errorf("unexpected panic for input %q", tt.input)
					}
				} else {
					if tt.shouldPanic {
						t.Errorf("expected panic for input %q, but did not panic", tt.input)
					}
				}
			}()

			actual := StringToUint64(tt.input)
			if actual != tt.expected && !tt.shouldPanic {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

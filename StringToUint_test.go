package convert

import (
	"testing"
)

func TestStringToUint(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    uint
		shouldPanic bool
	}{
		{
			name:        "Valid small number",
			input:       "123",
			expected:    123,
			shouldPanic: false,
		},
		{
			name:        "Valid large number within uint32 range",
			input:       "4294967295", // Max uint32
			expected:    4294967295,
			shouldPanic: false,
		},
		{
			name:        "Number exceeding uint32 range",
			input:       "4294967296", // Max uint32 + 1
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

			actual := StringToUint(tt.input)
			if actual != tt.expected && !tt.shouldPanic {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

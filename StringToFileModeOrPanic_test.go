package convert

import (
	"os"
	"testing"
)

func TestStringToFileModeOrPanic(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    os.FileMode
		shouldPanic bool
	}{
		{
			name:        "Valid file mode 0600",
			input:       "0600",
			expected:    0600,
			shouldPanic: false,
		},
		{
			name:        "Valid file mode 0777",
			input:       "0777",
			expected:    0777,
			shouldPanic: false,
		},
		{
			name:        "Valid file mode 1777",
			input:       "1777",
			expected:    01777,
			shouldPanic: false,
		},
		{
			name:        "Invalid file mode",
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

			actual := StringToFileModeOrPanic(tt.input)
			if actual != tt.expected && !tt.shouldPanic {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

package convert

import (
	"testing"
)

func TestMapToString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected string
	}{
		{
			name: "Non-empty map",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
				"key3": true,
			},
			expected: "key1=value1&key2=123&key3=true",
		},
		{
			name:     "Empty map",
			input:    map[string]interface{}{},
			expected: "",
		},
		{
			name:     "Nil map",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MapToString(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

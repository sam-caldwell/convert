package convert

import (
	"encoding/binary"
	"testing"
)

func TestUint16ToBigEndian(t *testing.T) {
	// Happy path: Little-endian input and conversion to big-endian
	tests := []struct {
		name     string
		input    uint16
		expected uint16
		endian   binary.ByteOrder
	}{
		{
			name:     "Convert little-endian 0x1234 to big-endian",
			input:    0x1234,
			expected: 0x3412,
			endian:   binary.LittleEndian,
		},
		{
			name:     "Already big-endian 0x3412",
			input:    0x3412,
			expected: 0x3412,
			endian:   binary.BigEndian,
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the native endian for testing
			NativeEndian = tt.endian
			result := Uint16ToBigEndian(tt.input)
			if result != tt.expected {
				t.Errorf("Uint16ToBigEndian(%#x) = %#x, want %#x", tt.input, result, tt.expected)
			}
		})
	}
}

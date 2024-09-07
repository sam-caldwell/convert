package convert

import (
	"encoding/binary"
	"testing"
)

func TestUint32ToBigEndian(t *testing.T) {
	// Happy path: Little-endian input and conversion to big-endian
	tests := []struct {
		name     string
		input    uint32
		expected uint32
		endian   binary.ByteOrder
	}{
		{
			name:     "Convert little-endian 0x12345678 to big-endian",
			input:    0x12345678,
			expected: 0x78563412,
			endian:   binary.LittleEndian,
		},
		{
			name:     "Already big-endian 0x78563412",
			input:    0x78563412,
			expected: 0x78563412,
			endian:   binary.BigEndian,
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the native endian for testing
			NativeEndian = tt.endian
			result := Uint32ToBigEndian(tt.input)
			if result != tt.expected {
				t.Errorf("Uint32ToBigEndian(%#x) = %#x, want %#x", tt.input, result, tt.expected)
			}
		})
	}
}

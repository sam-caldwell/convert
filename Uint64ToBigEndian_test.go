package convert

import (
	"encoding/binary"
	"testing"
)

func TestUint64ToBigEndian(t *testing.T) {
	// Happy path: Little-endian input and conversion to big-endian
	tests := []struct {
		name     string
		input    uint64
		expected uint64
		endian   binary.ByteOrder
	}{
		{
			name:     "Convert little-endian 0x123456789ABCDEF0 to big-endian",
			input:    0x123456789ABCDEF0,
			expected: 0xF0DEBC9A78563412,
			endian:   binary.LittleEndian,
		},
		{
			name:     "Already big-endian 0xF0DEBC9A78563412",
			input:    0xF0DEBC9A78563412,
			expected: 0xF0DEBC9A78563412,
			endian:   binary.BigEndian,
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the native endian for testing
			NativeEndian = tt.endian
			result := Uint64ToBigEndian(tt.input)
			if result != tt.expected {
				t.Errorf("Uint64ToBigEndian(%#x) = %#x, want %#x", tt.input, result, tt.expected)
			}
		})
	}
}

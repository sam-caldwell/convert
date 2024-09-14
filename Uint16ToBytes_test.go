package convert

import (
	"encoding/binary"
	"testing"
)

func TestUint16ToBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    uint16
		expected []byte
	}{
		{
			name:     "Zero value",
			input:    0,
			expected: []byte{0x00, 0x00},
		},
		{
			name:     "Small value",
			input:    1,
			expected: []byte{0x01, 0x00},
		},
		{
			name:     "Mid value1",
			input:    255,
			expected: []byte{0xFF, 0x00},
		},
		{
			name:     "Mid value2",
			input:    256,
			expected: []byte{0x00, 0x01},
		},
		{
			name:     "Max value",
			input:    65535,
			expected: []byte{0xFF, 0xFF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint16ToBytes(tt.input)
			if binary.NativeEndian.Uint16(result) != binary.NativeEndian.Uint16(tt.expected) {
				t.Errorf("Uint16ToBytes(%d) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

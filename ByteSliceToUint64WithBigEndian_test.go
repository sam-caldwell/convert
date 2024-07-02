package convert

import (
	"testing"
)

func TestByteSliceToUint64WithBigEndian(t *testing.T) {
	// Happy path
	happyPathInput := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	expectedOutput := uint64(1)
	actualOutput := ByteSliceToUint64WithBigEndian(happyPathInput)
	if actualOutput != expectedOutput {
		t.Errorf("expected %d, got %d", expectedOutput, actualOutput)
	}

	// Sad path: input length less than 8 bytes
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for input length less than 8 bytes, but did not panic")
		}
	}()
	ByteSliceToUint64WithBigEndian([]byte{0x00, 0x00, 0x00, 0x00})

	// Sad path: input length greater than 8 bytes
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for input length greater than 8 bytes, but did not panic")
		}
	}()
	ByteSliceToUint64WithBigEndian([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02})
}

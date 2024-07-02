package convert

import (
	"testing"
)

func TestByteToHexString(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		happyPathInput := []byte{0x0a, 0x1b, 0x2c, 0x3d, 0x4e, 0x5f}
		expectedOutput := "0a 1b 2c 3d 4e 5f "
		actualOutput := ByteToHexString(happyPathInput)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}
	})
	t.Run("Edge case: empty byte slice", func(t *testing.T) {
		var emptyInput []byte
		expectedOutputEmpty := ""
		actualOutputEmpty := ByteToHexString(emptyInput)
		if actualOutputEmpty != expectedOutputEmpty {
			t.Errorf("expected %q, got %q", expectedOutputEmpty, actualOutputEmpty)
		}
	})
	t.Run("Edge case: single byte", func(t *testing.T) {
		singleByteInput := []byte{0xff}
		expectedOutputSingleByte := "ff "
		actualOutputSingleByte := ByteToHexString(singleByteInput)
		if actualOutputSingleByte != expectedOutputSingleByte {
			t.Errorf("expected %q, got %q", expectedOutputSingleByte, actualOutputSingleByte)
		}
	})
}

package convert

import (
	"github.com/sam-caldwell/errors/v2"
	"testing"
)

func TestBytesToInt32(t *testing.T) {
	// Happy path
	happyPathInput := []byte{0x00, 0x00, 0x00, 0x01}
	expectedOutput := int32(1)
	actualOutput, err := BytesToInt32(happyPathInput)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if actualOutput != expectedOutput {
		t.Errorf("expected %d, got %d", expectedOutput, actualOutput)
	}

	// Sad path: nil input
	_, err = BytesToInt32(nil)
	if err == nil {
		t.Errorf("expected error for nil input, but got none")
	} else if err.Error() != errors.BoundsCheckError {
		t.Errorf("expected BoundsCheckError, but got %v", err)
	}

	// Sad path: input length less than 4 bytes
	_, err = BytesToInt32([]byte{0x00, 0x00, 0x00})
	if err == nil {
		t.Errorf("expected error for input length less than 4 bytes, but got none")
	} else if err.Error() != errors.BoundsCheckError {
		t.Errorf("expected BoundsCheckError, but got %v", err)
	}

	// Sad path: input length greater than 4 bytes
	_, err = BytesToInt32([]byte{0x00, 0x00, 0x00, 0x01, 0x02})
	if err == nil {
		t.Errorf("expected error for input length greater than 4 bytes, but got none")
	} else if err.Error() != errors.BoundsCheckError {
		t.Errorf("expected BoundsCheckError, but got %v", err)
	}
}

func main() {
	// Run tests
	t := &testing.T{}
	TestBytesToInt32(t)
}

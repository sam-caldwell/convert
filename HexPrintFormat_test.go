package convert

import (
	"bytes"
	"testing"
)

func TestHexPrintFormat(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		input := []byte{0x0a, 0x1b, 0x2c, 0x3d, 0x4e, 0x5f, 0x6a, 0x7b, 0x8c, 0x9d}
		width := 4
		output := "0a 1b 2c 3d \n4e 5f 6a 7b \n8c 9d \n"
		actual, err := HexPrintFormat(input, width)
		if err != nil {
			t.Fatalf("unexpected error: '%v'", err)
		}
		if actual != output {
			t.Fatalf("expected %q, got %q", output, actual)
		}
	})

	t.Run("Edge case: nil input", func(t *testing.T) {
		input := []byte(nil)
		width := 1337
		output := ""
		expectedError := errUnexpectedNil
		actual, err := HexPrintFormat(input, width)
		if err == nil {
			t.Fatalf("expected error not found")
		}
		if err != nil && err.Error() != expectedError {
			t.Fatalf("unexpected error %v", err)
		}
		if actual != output {
			t.Fatal("expected output not found")
		}
	})

	t.Run("Edge case: empty byte slice", func(t *testing.T) {
		var input []byte
		width := 4
		output := ""
		expectedError := errUnexpectedNil
		actual, err := HexPrintFormat(input, width)
		if err == nil {
			t.Fatal("expected error not found")
		}
		if err != nil && err.Error() != expectedError {
			t.Fatalf("unexpected error: '%v'", err)
		}
		if actual != output {
			t.Errorf("expected %q, got %q", output, actual)
		}
	})

	t.Run("Edge case: width too small", func(t *testing.T) {
		input := []byte{01, 02, 03}
		width := 0
		output := ""
		actual, err := HexPrintFormat(input, width)
		if err == nil {
			t.Fatalf("expected error (width too small)")
		}
		if err != nil && err.Error() != errMinimumWidth {
			t.Fatalf("expected minimum width error but got '%v'", err)
		}
		if actual != output {
			t.Fatalf("expected output not found\n"+
				"actual:   '%s'"+
				"expected: '%s'",
				actual, output)
		}
	})

	t.Run("Edge case: width of 1", func(t *testing.T) {
		input := []byte{0x0a, 0x1b, 0x2c}
		width := 1
		output := "0a \n1b \n2c \n"
		actual, err := HexPrintFormat(input, width)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if actual != output {
			t.Errorf("expected %q, got %q", output, actual)
		}
	})

	t.Run("Edge case: width greater than input length", func(t *testing.T) {
		input := []byte("test string")
		width := 90
		output := "74 65 73 74 20 73 74 72 69 6e 67 \n"
		actual, err := HexPrintFormat(input, width)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(actual) != len(output) {
			t.Errorf("size mismatch \n"+
				"actual: %d\n"+
				"output: %d\n",
				len(actual), len(output))
		}
		if !bytes.Equal([]byte(actual), []byte(output)) {
			t.Errorf("output mismatch:\n"+
				"expected '%s'\n"+
				"     got '%s'", output, actual)
		}
	})
}

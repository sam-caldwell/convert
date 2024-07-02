package convert

import (
	"testing"
)

func TestCapitalizep(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		Input := "hello"
		expectedOutput := "Hello"
		actualOutput := Capitalizep(&Input)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}
	})

	t.Run("Edge case: empty string", func(t *testing.T) {
		Input := ""
		expectedOutput := ""
		actualOutput := Capitalizep(&Input)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}

	})

	t.Run("Edge case: nil pointer", func(t *testing.T) {
		var Input *string
		expectedOutput := ""
		actualOutput := Capitalizep(Input)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}

	})

	t.Run("Edge case: string with first character already capitalized", func(t *testing.T) {
		Input := "Hello"
		expectedOutput := "Hello"
		actualOutput := Capitalizep(&Input)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}
	})

	t.Run("Edge case: string with non-alphabetic first character", func(t *testing.T) {
		Input := "123hello"
		expectedOutput := "123hello"
		actualOutput := Capitalizep(&Input)
		if actualOutput != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, actualOutput)
		}
	})
}

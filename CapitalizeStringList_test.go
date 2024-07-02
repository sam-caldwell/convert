package convert

import (
	"testing"
)

func TestCapitalizeStringList(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		// Happy path
		Input := []string{"hello", "world", "golang"}
		Output := []string{"Hello", "World", "Golang"}
		CapitalizeStringList(&Input)
		for i, v := range Input {
			if v != Output[i] {
				t.Errorf("expected %q, got %q at index %d", Output[i], v, i)
			}
		}
	})
	t.Run("Edge case: empty slice", func(t *testing.T) {
		var (
			Input  []string
			Output []string
		)
		CapitalizeStringList(&Input)
		if len(Input) != len(Output) {
			t.Errorf("expected empty slice, got %v", Input)
		}
	})
	t.Run("Edge case: slice with empty strings", func(t *testing.T) {
		Input := []string{"", "", ""}
		Output := []string{"", "", ""}
		CapitalizeStringList(&Input)
		for i, v := range Input {
			if v != Output[i] {
				t.Errorf("expected %q, got %q at index %d", Output[i], v, i)
			}
		}
	})
	t.Run("Edge case: slice with mixed cases", func(t *testing.T) {
		Input := []string{"hello", "World", "gOlAnG"}
		Output := []string{"Hello", "World", "GOlAnG"}
		CapitalizeStringList(&Input)
		for i, v := range Input {
			if v != Output[i] {
				t.Errorf("expected %q, got %q at index %d", Output[i], v, i)
			}
		}
	})
	t.Run("Edge case: slice with non-alphabetic first characters", func(t *testing.T) {
		Input := []string{"123hello", "!world", "golang", "_golang"}
		Output := []string{"123hello", "!world", "Golang", "_golang"}
		CapitalizeStringList(&Input)
		for i, v := range Input {
			if v != Output[i] {
				t.Errorf("expected %q, got %q at index %d", Output[i], v, i)
			}
		}
	})
}

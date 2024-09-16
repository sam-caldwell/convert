package convert

import "testing"

func TestBytesToUint16Be(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		if v := BytesToUint16Be([2]byte{0x00}); v != 0x00 {
			t.Fatal("failed with 0x00")
		}
		if v := BytesToUint16Be([2]byte{0x00, 0x01}); v != 0x0001 {
			t.Fatal("failed with 0x01")
		}
		if v := BytesToUint16Be([2]byte{0x10, 0x00}); v != 0x1000 {
			t.Fatal("failed with 0x1000")
		}
	})
}

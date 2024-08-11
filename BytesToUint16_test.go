package convert

import "testing"

func TestBytesToUint16(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		if v := BytesToUint16([2]byte{0x00}); v != 0x00 {
			t.Fatal("failed with 0x00")
		}
		if v := BytesToUint16([2]byte{0x00, 0x01}); v != 0x0001 {
			t.Fatal("failed with 0x01")
		}
		if v := BytesToUint16([2]byte{0x10, 0x00}); v != 0x1000 {
			t.Fatal("failed with 0x1000")
		}
	})
	t.Run("test range", func(t *testing.T) {
		for msb := 0; msb < 255; msb++ {
			for lsb := 0; lsb < 255; lsb++ {
				c := 256*msb + lsb
				if v := BytesToUint16([2]byte{byte(msb), byte(lsb)}); v != uint16(c) {
					t.Fatalf("Faild at %d:%d == %d", msb, lsb, c)
				}

			}
		}
	})
}

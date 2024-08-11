package convert

import "testing"

func TestBytesToUint32(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		if v := BytesToUint32([4]byte{0x00}); v != 0x00 {
			t.Fatal("failed with 0x00")
		}
		if v := BytesToUint32([4]byte{0x00, 0x00, 0x00, 0x01}); v != 0x0001 {
			t.Fatal("failed with 0x01")
		}
		if v := BytesToUint32([4]byte{0x00, 0x00, 0x10, 0x00}); v != 0x1000 {
			t.Fatal("failed with 0x1000")
		}
	})
	t.Run("test range", func(t *testing.T) {
		const (
			p0 = 256 * 256 * 256
			p1 = 256 * 256
			p2 = 256
			p3 = 1
		)
		const stagger = 64
		for b0 := 0; b0 < 256; b0 += stagger {
			for b1 := 0; b1 < 256; b1 += stagger {
				for b2 := 0; b2 < 256; b2 += stagger {
					for b3 := 0; b3 < 256; b3 += stagger {
						c := p0*b0 + p1*b1 + p2*b2 + p3*b3
						if v := BytesToUint32([4]byte{byte(b0), byte(b1), byte(b2), byte(b3)}); v != uint32(c) {
							t.Fatalf("Failed at %d:%d:%d:%d == %d", b0, b1, b2, b3, c)
						}
					}
				}
			}
		}
	})
}

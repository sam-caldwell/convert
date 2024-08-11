package convert

import "testing"

func TestBytesToUint64(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		if v := BytesToUint64([8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}); v != 0x00 {
			t.Fatal("failed with 0x00")
		}
		if v := BytesToUint64([8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}); v != 0x0001 {
			t.Fatal("failed with 0x01")
		}
		if v := BytesToUint64([8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00}); v != 0x1000 {
			t.Fatal("failed with 0x1000")
		}
	})
	t.Run("test range", func(t *testing.T) {
		const (
			p0 = 256 * 256 * 256 * 256 * 256 * 256 * 256
			p1 = 256 * 256 * 256 * 256 * 256 * 256
			p2 = 256 * 256 * 256 * 256 * 256
			p3 = 256 * 256 * 256 * 256
			p4 = 256 * 256 * 256
			p5 = 256 * 256
			p6 = 256
			p7 = 1
		)
		const stagger = 127
		for b0 := 0; b0 < 256; b0 += stagger {
			for b1 := 0; b1 < 256; b1 += stagger {
				for b2 := 0; b2 < 256; b2 += stagger {
					for b3 := 0; b3 < 256; b3 += stagger {
						for b4 := 0; b4 < 256; b4 += stagger {
							for b5 := 0; b5 < 256; b5 += stagger {
								for b6 := 0; b6 < 256; b6 += stagger {
									for b7 := 0; b7 < 256; b7 += stagger {
										c := p0*b0 + p1*b1 + p2*b2 + p3*b3 + p4*b4 + p5*b5 + p6*b6 + p7*b7
										v := BytesToUint64([8]byte{
											byte(b0), byte(b1), byte(b2), byte(b3),
											byte(b4), byte(b5), byte(b6), byte(b7),
										})
										if v != uint64(c) {
											t.Fatalf("Failed at %d:%d:%d:%d:%d:%d:%d:%d == %d",
												b0, b1, b2, b3, b4, b5, b6, b7, c)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	})
}

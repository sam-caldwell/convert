package convert

import (
	"testing"
)

func TestBytesToKilobytes(t *testing.T) {
	data := map[int]int{
		1:       0,
		2:       0,
		512:     0,
		1023:    0,
		1024:    1,
		1500:    1,
		2048:    2,
		3000:    2,
		3072:    3,
		4096:    4,
		8192:    8,
		16384:   16,
		32768:   32,
		65536:   64,
		1048576: 1024,
	}
	for b, kb := range data {
		if actual := BytesToKilobytes(b); actual != kb {
			t.Fatalf("BytesToKilobytes(%d) expected %d but returned %d", b, kb, actual)
		}
	}
}

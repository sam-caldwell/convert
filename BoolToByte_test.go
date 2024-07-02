package convert

import (
	"testing"
)

func TestBoolToByte(t *testing.T) {
	t.Run("true bool to 0x01", func(t *testing.T) {
		if o := BoolToByte(true); o != 0x01 {
			t.Fatalf("BoolToByte(true) expects 0x01")
		}
	})
	t.Run("false bool to 0x00", func(t *testing.T) {
		if o := BoolToByte(false); o != 0x00 {
			t.Fatalf("BoolToByte(false) expects 0x00")
		}
	})

}

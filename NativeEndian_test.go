package convert

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestNativeEndianInitialization(t *testing.T) {
	var x uint32 = 0x01020304
	var expectedEndian binary.ByteOrder

	if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
		expectedEndian = binary.BigEndian
	} else {
		expectedEndian = binary.LittleEndian
	}

	if NativeEndian != expectedEndian {
		t.Errorf("NativeEndian was initialized incorrectly. Got %v, expected %v", NativeEndian, expectedEndian)
	}
}

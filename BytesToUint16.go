package convert

import "encoding/binary"

// BytesToUint16 - Convert [2]byte into uint16 with NativeEndian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint16(b [2]byte) uint16 {
	return binary.NativeEndian.Uint16(b[:])
}

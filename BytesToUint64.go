package convert

import "encoding/binary"

// BytesToUint64 - Convert [8]byte into uint16 with NativeEndian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint64(b [8]byte) uint64 {
	return binary.NativeEndian.Uint64(b[:])
}

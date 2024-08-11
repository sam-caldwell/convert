package convert

import "encoding/binary"

// BytesToUint32 - Convert [4]byte into uint16 with BigEndian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint32(b [4]byte) uint32 {
	return binary.BigEndian.Uint32(b[:])
}

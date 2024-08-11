package convert

import "encoding/binary"

// BytesToUint64 - Convert [8]byte into uint16 with BigEndian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint64(b [8]byte) uint64 {
	return binary.BigEndian.Uint64(b[:])
}

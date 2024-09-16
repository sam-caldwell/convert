package convert

import "encoding/binary"

// BytesToUint64Be - Convert [8]byte into uint16 with Big Endian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint64Be(b [8]byte) uint64 {
	return binary.BigEndian.Uint64(b[:])
}

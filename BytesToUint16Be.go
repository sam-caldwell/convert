package convert

import "encoding/binary"

// BytesToUint16Be - Convert [2]byte into uint16 with Big Endian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint16Be(b [2]byte) uint16 {
	return binary.BigEndian.Uint16(b[:])
}

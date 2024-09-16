package convert

import "encoding/binary"

// BytesToUint32Be - Convert [4]byte into uint16 with Big Endian.
//
//	This is particularly useful when parsing Ethernet frames.
//	Have fun!
func BytesToUint32Be(b [4]byte) uint32 {
	return binary.BigEndian.Uint32(b[:])
}

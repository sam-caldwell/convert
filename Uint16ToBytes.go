package convert

import "encoding/binary"

// Uint16ToBytes - Convert uint16 to []byte
func Uint16ToBytes(value uint16) []byte {
	bytes := make([]byte, 2)
	binary.NativeEndian.PutUint16(bytes, value)
	return bytes
}

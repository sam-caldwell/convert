package convert

import "fmt"

// ByteToHexString - convert a []byte array to a hex string.
func ByteToHexString(b []byte) (out string) {
	for i := 0; i < len(b); i++ {
		out += fmt.Sprintf("%02x ", b[i])
	}
	return out
}

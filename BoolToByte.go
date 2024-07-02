package convert

// BoolToByte - Convert Boolean true/false to a byte value
func BoolToByte(b bool) byte {
	if b {
		return 0x01
	}
	return 0x00
}

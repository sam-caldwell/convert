package convert

import (
	"fmt"
	"github.com/sam-caldwell/errors/v2"
)

// BytesToInt32 - Convert byte array to int32
func BytesToInt32(b []byte) (int32, error) {
	if b == nil || len(b) != 4 {
		return 0, fmt.Errorf(errors.BoundsCheckError)
	}
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3]), nil
}

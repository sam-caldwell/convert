package convert

import (
	"os"
	"strconv"
)

// StringToFileModeOrPanic - Convert a string permission (e.g. 0600) to FileMode numeric value
func StringToFileModeOrPanic(value string) os.FileMode {
	number, err := strconv.ParseUint(value, 8, 64)
	if err != nil {
		panic(err)
	}
	return os.FileMode(number)
}

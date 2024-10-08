package convert

import (
	"github.com/sam-caldwell/errors"
	"math"
	"strconv"
)

// StringToUint - Convert the string value to an unsigned-integer.
func StringToUint(value string) (n uint) {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(err)
	}
	if number <= math.MaxUint32 {
		return uint(number)
	}
	panic(errors.BoundsCheckError)
}

package convert

import "fmt"

const (
	errUnexpectedNil = "unexpected nil"
	errMinimumWidth  = "width must be >= 1"
)

// HexPrintFormat - Given []byte, return a hex-formatted string with a line feed (\n) character at every width position.
func HexPrintFormat(input []byte, width int) (out string, err error) {
	if input == nil {
		return "", fmt.Errorf(errUnexpectedNil)
	}
	if width < 1 {
		return "", fmt.Errorf(errMinimumWidth)
	}
	//if width >= len(input) {
	//	return string(input) + "\n", nil
	//}
	for pos, b := range input {
		if pos > 0 && pos%width == 0 {
			out += "\n"
		}
		out += fmt.Sprintf("%02x ", b)
	}
	if len(input) > 0 {
		out += "\n"
	}
	return out, nil
}

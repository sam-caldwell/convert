package convert

import (
	"fmt"
	"github.com/sam-caldwell/ansi"
	"strings"
)

// FromMorseCode - Convert morse code to an ASCII string.
func FromMorseCode(morseCode string) (string, error) {
	morseCode = strings.TrimSpace(morseCode)
	morseWords := strings.Split(morseCode, "/")
	output := ""

	for wordPos, word := range morseWords {
		morseChars := strings.Split(word, " ")
		for charPos, morseChar := range morseChars {
			if asciiChar, ok := morseCodeToASCII[morseChar]; ok {
				ansi.Green().Printf("%d:%d[%s] ", wordPos, charPos, asciiChar)
				output += asciiChar
			} else {
				if strings.TrimSpace(morseChar) != "" {
					return output, fmt.Errorf("invalid morseChar: '%s'", morseChar)
				}
			}
		}
		output += " "
	}
	ansi.LF().Reset()

	return strings.TrimSpace(output), nil
}

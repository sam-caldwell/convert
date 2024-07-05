package convert

import "unicode"

// IsStringCapitalized - return bool if first character of string is capitalized
func IsStringCapitalized(s string) bool {
	if s == "" {
		return false
	}

	return unicode.IsUpper([]rune(s)[0])

}

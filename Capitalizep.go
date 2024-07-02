package convert

import "unicode"

// Capitalizep - Capitalize the first character of the *string
func Capitalizep(s *string) string {
	if s == nil {
		return ""
	}
	if *s == "" {
		return *s
	}
	runes := []rune(*s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

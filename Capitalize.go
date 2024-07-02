package convert

// Capitalize - Capitalize the first character of the string
//
//go:inline
func Capitalize(s string) string {
	return Capitalizep(&s)
}

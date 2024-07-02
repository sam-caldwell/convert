package convert

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// ErrorToString - Safely turn an error into a string
func ErrorToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return words.EmptyString
}

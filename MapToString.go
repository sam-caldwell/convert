package convert

import (
	"fmt"
	"strings"
)

// MapToString - Given a map of generic types return a string of ampersand (&) delimited key=value pairs
//
//	This function will take a map[K]V and return a string, delimited by an ampersand (&)
//	For example:
//	   key1=value1&key2=value2&key3=value3
func MapToString[K string, V any](tags map[K]V) string {
	var tagPairs []string

	if tags == nil {
		return ""
	}

	for k, v := range tags {
		tagPairs = append(tagPairs, fmt.Sprintf("%v=%v", k, v))
	}

	return strings.Join(tagPairs, "&")

}

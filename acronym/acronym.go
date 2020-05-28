// Package acronym create acronyms from strings
package acronym

import (
	"unicode"
)

// Abbreviate a string into its acronym
func Abbreviate(s string) string {
	acronym := []rune{}
	nw := true
	for _, c := range s {
		if c == ' ' || c == '-' {
			nw = true
			continue
		}
		if nw == true && unicode.IsLetter(c) {
			acronym = append(acronym, unicode.ToUpper(c))
			nw = false
		}
	}
	return string(acronym)
}

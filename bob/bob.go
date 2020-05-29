// Package bob ... What About Bob?
package bob

import (
	"strings"
	"unicode"
)

// Hey responds to a remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	question := strings.HasSuffix(remark, "?")
	shouting := isShouting(remark)

	switch {
	case shouting && question:
		return "Calm down, I know what I'm doing!"
	case shouting:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	}

	return "Whatever."
}

func isShouting(s string) (shouting bool) {
	shouting = true
	hasLetters := false
	for _, c := range s {
		if shouting && unicode.IsLower(c) {
			shouting = false
		}
		if !hasLetters && unicode.IsLetter(c) {
			hasLetters = true
		}
	}
	if shouting && !hasLetters {
		shouting = false
	}
	return
}

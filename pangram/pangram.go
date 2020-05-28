package pangram

import (
	"unicode"
)

//IsPangram determins if phrase is a pangram
func IsPangram(phrase string) bool {
	count := make(map[rune]int)
	for _, c := range phrase {
		if unicode.IsLetter(c) {
			count[unicode.ToLower(c)]++
		}
	}
	return len(count) == 26
}

package luhn

import (
	"strings"
)

//Valid checks if input validates against Luhn https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(input string) bool {
	var sum int
	input = strings.ReplaceAll(input, " ", "")

	l := len(input)

	if l <= 1 {
		return false
	}

	dbl := l%2 == 0
	for _, r := range input {
		d := int(r - '0')
		if d < 0 || d > 9 {
			return false
		}
		if dbl {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		dbl = !dbl
	}
	return sum%10 == 0
}

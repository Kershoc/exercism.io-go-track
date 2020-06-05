package isbn

import (
	"strings"
)

//IsValidISBN determines if input is a valid ISBN-10 number https://en.wikipedia.org/wiki/International_Standard_Book_Number
func IsValidISBN(input string) bool {
	input = strings.ReplaceAll(input, "-", "")
	if len(input) != 10 {
		return false
	}
	sum := 0
	for i, r := range input {
		if i == 9 && r == 'X' {
			sum += 10
			continue
		}
		d := int(r - '0')
		if d < 0 || d > 9 {
			return false
		}
		sum += d * (10 - i)
	}
	return sum%11 == 0
}

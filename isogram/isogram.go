package isogram

import "unicode"

//IsIsogram returns true when there are no repeated letters in the word
func IsIsogram(word string) bool {
	uniqueRunes := make(map[rune]bool)
	var length int

	for _, character := range word {
		if unicode.IsLetter(character) {
			uniqueRunes[unicode.ToLower(character)] = true
			length++
		}
	}

	return len(uniqueRunes) == length
}

package etl

import "strings"

//Transform translates scrabble scores stored in format map[score][]letters to map[letter]score
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}

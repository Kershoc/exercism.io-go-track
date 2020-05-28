package anagram

import (
	"sort"
	"strings"
)

//Detect will return all anagrams of input found in candidates
func Detect(input string, candidates []string) []string {
	output := []string{}
	input = strings.ToLower(input)
	sortedInput := sortString(input)
	for _, candidate := range candidates {
		check := strings.ToLower(candidate)
		if sortedInput == sortString(check) && input != check {
			output = append(output, candidate)
		}
	}
	return output
}

func sortString(input string) string {
	letters := strings.Split(input, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

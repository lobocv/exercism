package isogram

import "strings"

func IsIsogram(word string) bool {
	// Make all letters the same case (upper)
	word = strings.ToUpper(word)

	// Remove any hyphens or spaces so we are only left with characters
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)

	// Create a map that we can push letters into
	var letterdict = make(map[rune]bool)

	// Push letters into the map one at a time, if it already exists in the map then it's a duplicate
	for _, letter := range word{
		if _, ok := letterdict[letter]; ok {
			return false
		}
		letterdict[letter] = true
	}
	return true
}
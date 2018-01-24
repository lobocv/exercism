// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"bytes"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Return the abbreviation for the argument string.

	// Create a buffer we can place acronym letters in
	var acronym bytes.Buffer

	// Split the string into words and then split each word by hyphens
	// Place the capitalized first letter into the buffer
	for _, word := range strings.Split(s, " ") {
		for _, subword := range strings.Split(word, "-") {
			firstLetter := string(subword[0])
			acronym.WriteString(strings.ToUpper(firstLetter))
		}
	}

	return acronym.String()
}

// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isShouting(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ContainsAny(remark, alphabet)
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}


// Hey should have a comment documenting it.
func Hey(remark string) string {

	// Strip the remark of any blank characters such as new lines and tabs
	remark = strings.Trim(remark, " \t\r\n")

	if isShouting(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark){
		return "Sure."
	} else if isShouting(remark){
		return "Whoa, chill out!"
	} else if strings.EqualFold(remark,""){
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

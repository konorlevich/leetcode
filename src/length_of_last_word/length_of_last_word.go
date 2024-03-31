// Package length_of_last_word
//
// Given a string s consisting of words and spaces,
// return the length of the last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.
package length_of_last_word

import "strings"

func lengthOfLastWord(s string) int {
	lastIndex := max(len(strings.TrimRight(s, " "))-1, 0)
	for i := lastIndex; i >= 0; i-- {
		if s[i] == ' ' {
			return lastIndex - i
		}
	}

	return lastIndex + 1
}

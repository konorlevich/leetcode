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

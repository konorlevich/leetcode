// Package longest_palindromic_substring
//
// Given a string s, return the longest palindromic substring in s.
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters.
package longest_palindromic_substring

// longestPalindrome a straightforward n^2 solution
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// although we promised to receive only ascii, it's easier to iterate over rune set
	sAsRune := []rune(s)

	longest := make([]rune, 0, len(sAsRune))

	// iterate over the rune list, using each symbol as a center
	for c := 0; c < len(sAsRune)-1; c++ {
		// left and right `window` set to the center
		l, r := c, c
		//in some cases we can have either `aa` or `aaa` symmetry
		for r < len(sAsRune)-1 && sAsRune[r+1] == sAsRune[l] {
			r++
		}
		// then just widen the `window`, checking symbols
		// save longest found palindrome
		for l >= 0 && r < len(sAsRune) {
			if sAsRune[l] != sAsRune[r] {
				break
			}
			if len(longest) == 0 && l == r {
				longest = append(longest, sAsRune[l])
			} else if r-l+1 > len(longest) {
				longest = sAsRune[l : r+1]
			}
			l--
			r++
		}
	}

	return string(longest)
}

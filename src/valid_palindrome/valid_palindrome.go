// Package valid_palindrome
//
// A phrase is a palindrome if,
// after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// s consists only of printable ASCII characters.
package valid_palindrome

const (
	ascii0              = uint8(48)
	ascii9              = uint8(57)
	asciiLetterACapital = uint8(65)
	asciiLetterZCapital = uint8(90)
	asciiLetterA        = uint8(97)
	asciiLetterZ        = uint8(122)

	distance = asciiLetterA - asciiLetterACapital
)

// isPalindrome
//
// Reads symbol by symbol from both sides,
// checks if both symbols are alphanumeric,
// checks if the symbols are equal
//
// s consists only of printable ASCII characters.
func isPalindrome(s string) bool {
	leftI := 0
	rightI := len(s) - 1

	for leftI <= rightI {
		if !isAlphanumeric(s[leftI]) {
			leftI++
			continue
		}
		if !isAlphanumeric(s[rightI]) {
			rightI--
			continue
		}
		if areSimilar(s[leftI], s[rightI]) {
			leftI++
			rightI--
		} else {
			return false
		}
	}
	return true
}

// areSimilar checks if two ASCII-symbols are equal (case-insensitive)
func areSimilar(s1, s2 uint8) bool {
	if s1 == s2 {
		return true
	}
	if s1 > asciiLetterZCapital {
		s1 -= distance
	}
	if s2 > asciiLetterZCapital {
		s2 -= distance
	}
	if s1 != s2 {
		return false
	}
	return true
}

// isAlphanumeric checks if ASCII-symbol is 0-9 or a-z or A-Z
func isAlphanumeric(s uint8) bool {
	if s < ascii0 {
		return false
	}
	if s > ascii9 && s < asciiLetterACapital {
		return false
	}
	if s > asciiLetterZCapital && s < asciiLetterA {
		return false
	}
	if s > asciiLetterZ {
		return false
	}
	return true
}

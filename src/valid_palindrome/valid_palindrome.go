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

import "github.com/konorlevich/leetcode/src/common/ascii"

const (
	distance = ascii.LetterA - ascii.LetterACapital
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
	if s1 > ascii.LetterZCapital {
		s1 -= distance
	}
	if s2 > ascii.LetterZCapital {
		s2 -= distance
	}
	if s1 != s2 {
		return false
	}
	return true
}

// isAlphanumeric checks if ASCII-symbol is 0-9 or a-z or A-Z
func isAlphanumeric(s uint8) bool {
	if s < ascii.Digit0 {
		return false
	}
	if s > ascii.Digit9 && s < ascii.LetterACapital {
		return false
	}
	if s > ascii.LetterZCapital && s < ascii.LetterA {
		return false
	}
	if s > ascii.LetterZ {
		return false
	}
	return true
}

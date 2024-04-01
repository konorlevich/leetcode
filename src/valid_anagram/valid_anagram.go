// Package valid_anagram
//
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
package valid_anagram

import "github.com/konorlevich/leetcode/src/common"

// iterateWithMap
//
// We iterate over both strings, saving letters as a counter to two maps.
// Then we check lengths of maps.
// And then we iterate over one of the maps, and checking if in both maps counters are equal.
func iterateWithMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letters := make(map[string]map[uint8]int, 2)
	letters[s], letters[t] = make(map[uint8]int, len(s)), make(map[uint8]int, len(t))

	for i := 0; i < len(s); i++ {
		letters[s][s[i]] = letters[s][s[i]] + 1
		letters[t][t[i]] = letters[t][t[i]] + 1
	}
	if len(letters[s]) != len(letters[t]) {
		return false
	}
	for u, i := range letters[s] {
		if letters[t][u] != i {
			return false
		}
	}
	return true
}

// sumOfPower
//
// We iterate over strings, summing numerical representation of letters to the power of 4.
// Then we check if sums are equal.
func sumOfPower(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	var sumS, sumT = 0, 0

	for i := 0; i < len(s); i++ {
		sumS += common.Pow(int(s[i]), 4)
		sumT += common.Pow(int(t[i]), 4)
	}
	return sumS == sumT
}

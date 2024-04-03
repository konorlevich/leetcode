// Package isomorphic_strings
//
// Given two strings s and t, determine if they are isomorphic.
//
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to itself.
//
// # Constrains
//
// 1 <= s.length <= 5 * 104
//
// t.length == s.length
//
// s and t consist of any valid ascii character.
package isomorphic_strings

// withMap
//
// We iterate over literals in s, looking if we had a replacement for it in t.
// Then we check if we had a replacement for the corresponding literal in t.
// Then we save both replacements
func withMap(s, t string) bool {
	if s == t {
		return true
	}

	replaces := make(map[uint8]uint8, len(s))
	replacesReversed := make(map[uint8]uint8, len(s))
	for i := 0; i < len(s); i++ {
		sl := s[i]
		tl := t[i]
		if st, ok := replaces[sl]; ok {
			if tl != st {
				return false
			}
			continue
		}
		if ts, ok := replacesReversed[tl]; ok {
			if sl != ts {
				return false
			}
			continue
		}

		replaces[sl] = tl
		replacesReversed[tl] = sl
	}
	return true
}

// withArray
//
// The solution I found out there.
//
// We create two arrays of counters of integer representation of string literals for both strings.
// Then we iterate over the string literals, comparing how many each literal we saw.
// If both literals used equally, we increment the counters.
// Return false otherwise.
func withArray(s, t string) bool {
	var literalUsedInS, literalUsedInT = [256]int{}, [256]int{}
	for i := range s {
		if literalUsedInS[s[i]] != literalUsedInT[t[i]] {
			return false
		}
		literalUsedInS[s[i]] = i + 1
		literalUsedInT[t[i]] = i + 1
	}
	return true
}

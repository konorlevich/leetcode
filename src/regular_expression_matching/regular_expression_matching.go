// Package regular_expression_matching
//
// Given an input string s and a pattern p,
// implement regular expression matching
// with support for '.' and '*' where:
//
// '.' Matches any single character.
//
// '*' Matches zero or more of the preceding element.
//
// The matching should cover the entire input string (not partial).
//
// Constraints:
//
// 1 <= s.length <= 20
//
// 1 <= p.length <= 20
//
// s contains only lowercase English letters.
//
// p contains only lowercase English letters, '.', and '*'.
//
// It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
package regular_expression_matching

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	lenS := len(s)
	lenP := len(p)

	if lenP == 0 {
		return lenS == 0
	}

	if lenP == 1 {
		if lenS != 1 {
			return false
		}
		if p == s {
			return true
		}
		if p == "." {
			return true
		}
		return false
	}
	if lenP == 2 {
		if p[1] != '*' {
			if len(s) < 2 {
				return false
			}
			return isMatch(s[0:1], p[0:1]) && isMatch(s[1:], p[1:])
		}

		if lenS == 0 {
			return true
		}

		if p[0] == '.' {
			return true
		}
		if p[0] != s[0] {
			return false
		}
		if lenS == 1 {
			return true
		}

		for i := 1; i < len(s); i++ {
			if s[i] != s[0] {
				return false
			}
		}
		return true
	}

	pChunkR := 1
	if p[1] == '*' {
		pChunkR = 2
	}
	pChunk := p[0:pChunkR]
	for i := lenS; i >= 0; i-- {
		var sChunk string
		if i == lenS {
			sChunk = s[0:]
		} else {
			sChunk = s[0:i]
		}
		if isMatch(sChunk, pChunk) {
			sChunk := s[i:]
			pChunk := p[pChunkR:]
			if isMatch(sChunk, pChunk) {
				return true
			}
		}
	}
	return false
}

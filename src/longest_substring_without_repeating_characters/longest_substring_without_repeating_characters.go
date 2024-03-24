package longest_substring_without_repeating_characters

// lengthOfLongestSubstring
// Given a string s, find the length of the longest substring without repeating characters.
//
// A substring is a contiguous non-empty sequence of characters within a string.
func lengthOfLongestSubstring(s string) (maxLen int) {
	known := make([]int32, 0, len(s))
	curLen := 0

LOOP:
	for i, char := range s {
		known = append(known, char)
		if i == 0 {
			curLen++
			maxLen++
			continue
		}
		for i2 := i - 1; i2 >= i-curLen; i2-- {
			currentKnown := known[i2]
			if currentKnown == char {
				curLen = i - i2
				continue LOOP
			}
		}

		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

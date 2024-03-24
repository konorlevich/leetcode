package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "empty"},

		{name: "abcabcbb",
			s:    "abcabcbb",
			want: 3},

		{name: "bbbbb",
			s:    "bbbbb",
			want: 1},

		//The answer is "wke", with the length of 3.
		//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
		{name: "pwwkew",
			s:    "pwwkew",
			want: 3},

		{name: "one space",
			s:    " ",
			want: 1},

		{name: "dvdf",
			s:    "dvdf",
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

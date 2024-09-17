package longest_palindromic_substring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "empty", s: "", want: ""},
		{name: "a", s: "a", want: "a"},
		{name: "ab", s: "ab", want: "a"},
		{name: "bb", s: "bb", want: "bb"},
		{name: "aba", s: "aba", want: "aba"},
		{name: "sdfabba", s: "sdfabba", want: "abba"},
		{name: "sdfaba", s: "sdfaba", want: "aba"},
		{name: "abaasd", s: "abaasd", want: "aba"},
		{name: "abaaasd", s: "abaaasd", want: "aba"},
		{name: "abaaaasd", s: "abaaaasd", want: "aaaa"},
		{name: "babad", s: "babad", want: "bab"},
		{name: "cbbd", s: "cbbd", want: "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

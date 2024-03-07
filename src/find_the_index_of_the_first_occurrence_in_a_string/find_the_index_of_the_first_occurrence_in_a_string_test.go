package find_the_index_of_the_first_occurrence_in_a_string

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{name: "empty", want: -1},
		{name: "1", haystack: "sadbutsad", needle: "sad", want: 0},
		{name: "2", haystack: "leetcode", needle: "leeto", want: -1},
		{name: "3", haystack: "hello", needle: "ll", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

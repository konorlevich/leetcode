package longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{name: "empty"},
		{name: "1", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "2", strs: []string{"dog", "racecar", "car"}, want: ""},
		{name: "3", strs: []string{"a"}, want: "a"},
		{name: "4", strs: []string{"ab", "a"}, want: "a"},
		{name: "5", strs: []string{"abab", "aba", ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

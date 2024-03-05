package palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{name: "empty", x: 0, want: true},
		{name: "10", x: 10, want: false},
		{name: "101", x: 101, want: true},
		{name: "-121", x: -121, want: false},
		{name: "10", x: 10, want: false},
	}
	funcs := map[string]func(x int) bool{
		"asString": asString,
		"asNumber": asNumber,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := isPalindrome(f, tt.x); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

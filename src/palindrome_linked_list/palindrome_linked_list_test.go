package palindrome_linked_list

import (
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

var funcs = map[string]func(*list.Node) bool{
	"with buffer": withBuffer,
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head *list.Node
		want bool
	}{
		{name: "empty", want: true},

		{name: "[1]",
			head: list.Create([]int{1}),
			want: true},
		{name: "[1,2,2,1]",
			head: list.Create([]int{1, 2, 2, 1}),
			want: true},
		{name: "[1,2]",
			head: list.Create([]int{1, 2}),
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := f(tt.head); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

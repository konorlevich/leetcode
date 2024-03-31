package add_two_numbers

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *list.Node
		l2   *list.Node
		want *list.Node
	}{
		{name: "empty"},

		{name: "l1 = [2,4,3], l2 = [5,6,4]",
			l1:   list.Create([]int{2, 4, 3}),
			l2:   list.Create([]int{5, 6, 4}),
			want: list.Create([]int{7, 0, 8})},

		{name: "l1 = [0], l2 = [0]",
			l1:   list.Create([]int{0}),
			l2:   list.Create([]int{0}),
			want: list.Create([]int{0})},

		{name: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]",
			l1:   list.Create([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   list.Create([]int{9, 9, 9, 9}),
			want: list.Create([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(tt.want, addTwoNumbers(tt.l1, tt.l2)); diff != "" {
				t.Errorf("addTwoNumbers()\n%s", diff)
			}
		})
	}
}

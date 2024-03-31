package merge_two_sorted_lists

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *list.Node
		list2 *list.Node
		want  *list.Node
	}{
		{name: "empty"},
		{
			name:  "1",
			list1: list.Create([]int{1, 2, 3}),
			list2: list.Create([]int{4, 5, 6}),
			want:  list.Create([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "2",
			list1: list.Create([]int{1, 2, 4}),
			list2: list.Create([]int{3, 5, 6}),
			want:  list.Create([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "3",
			list1: list.Create([]int{1, 2, 4}),
			list2: list.Create([]int{1, 3, 4}),
			want:  list.Create([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "4",
			list1: list.Create(nil),
			list2: list.Create([]int{0}),
			want:  list.Create([]int{0}),
		},
		{
			name:  "5",
			list1: list.Create([]int{2}),
			list2: list.Create([]int{1}),
			want:  list.Create([]int{1, 2}),
		},
		{
			name:  "6",
			list1: list.Create([]int{5}),
			list2: list.Create([]int{1, 2, 4}),
			want:  list.Create([]int{1, 2, 4, 5}),
		},
		{
			name:  "7",
			list1: list.Create([]int{5, 6}),
			list2: list.Create([]int{2, 1, 4}),
			want:  list.Create([]int{1, 2, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			diff := cmp.Diff(tt.want, mergeTwoLists(tt.list1, tt.list2))
			if diff != "" {

				t.Errorf("mergeTwoLists()\n%s", diff)
			}
		})
	}
}

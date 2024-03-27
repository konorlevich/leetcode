package merge_two_sorted_lists

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *common.ListNode
		list2 *common.ListNode
		want  *common.ListNode
	}{
		{name: "empty"},
		{
			name:  "1",
			list1: common.CreateList([]int{1, 2, 3}),
			list2: common.CreateList([]int{4, 5, 6}),
			want:  common.CreateList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "2",
			list1: common.CreateList([]int{1, 2, 4}),
			list2: common.CreateList([]int{3, 5, 6}),
			want:  common.CreateList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "3",
			list1: common.CreateList([]int{1, 2, 4}),
			list2: common.CreateList([]int{1, 3, 4}),
			want:  common.CreateList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "4",
			list1: common.CreateList(nil),
			list2: common.CreateList([]int{0}),
			want:  common.CreateList([]int{0}),
		},
		{
			name:  "5",
			list1: common.CreateList([]int{2}),
			list2: common.CreateList([]int{1}),
			want:  common.CreateList([]int{1, 2}),
		},
		{
			name:  "6",
			list1: common.CreateList([]int{5}),
			list2: common.CreateList([]int{1, 2, 4}),
			want:  common.CreateList([]int{1, 2, 4, 5}),
		},
		{
			name:  "7",
			list1: common.CreateList([]int{5, 6}),
			list2: common.CreateList([]int{2, 1, 4}),
			want:  common.CreateList([]int{1, 2, 4, 5, 6}),
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

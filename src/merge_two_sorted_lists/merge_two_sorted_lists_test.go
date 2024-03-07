package merge_two_sorted_lists

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{name: "empty"},
		{
			name:  "1",
			list1: createList([]int{1, 2, 3}),
			list2: createList([]int{4, 5, 6}),
			want:  createList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "2",
			list1: createList([]int{1, 2, 4}),
			list2: createList([]int{3, 5, 6}),
			want:  createList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:  "3",
			list1: createList([]int{1, 2, 4}),
			list2: createList([]int{1, 3, 4}),
			want:  createList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "4",
			list1: createList(nil),
			list2: createList([]int{0}),
			want:  createList([]int{0}),
		},
		{
			name:  "5",
			list1: createList([]int{2}),
			list2: createList([]int{1}),
			want:  createList([]int{1, 2}),
		},
		{
			name:  "6",
			list1: createList([]int{5}),
			list2: createList([]int{1, 2, 4}),
			want:  createList([]int{1, 2, 4, 5}),
		},
		{
			name:  "7",
			list1: createList([]int{5, 6}),
			list2: createList([]int{2, 1, 4}),
			want:  createList([]int{1, 2, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := cmp.Diff(tt.want, mergeTwoLists(tt.list1, tt.list2))
			if diff != "" {

				t.Errorf("mergeTwoLists()\n%s", diff)
			}
		})
	}
}

func Test_createList(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *ListNode
	}{
		{name: "empty"},
		{
			name: "valid", nums: []int{1, 2, 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createList(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createList(nums []int) (l *ListNode) {
	curNode := &ListNode{}
	for _, i := range nums {
		if l == nil {
			curNode.Val = i
			l = curNode
			continue
		}
		curNode.Next = &ListNode{Val: i}
		curNode = curNode.Next
	}
	return l
}

package intersection_of_two_linked_lists

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	headA1, headB1, tail1 := common.CreateTwoIntersectedLists([]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3})
	headA2, headB2, tail2 := common.CreateTwoIntersectedLists([]int{1, 9, 1}, []int{3}, []int{2, 4})
	tests := []struct {
		name  string
		headA *common.ListNode
		headB *common.ListNode
		want  *common.ListNode
	}{
		{name: "empty"},

		{name: "no intersections",
			headA: common.CreateList([]int{1, 2, 3}),
			headB: common.CreateList([]int{1, 2, 3})},

		{name: "A=[1,2,3] B=[1,2,3] tail=[1,2,3]",
			headA: headA1,
			headB: headB1,
			want:  tail1,
		},

		{name: "A=[1,9,1] B=[3] tail=[2,4]",
			headA: headA2,
			headB: headB2,
			want:  tail2,
		},

		{name: "A=[2,6,4] B=[1,5]",
			headA: common.CreateList([]int{2, 6, 4}),
			headB: common.CreateList([]int{1, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, getIntersectionNode(tt.headA, tt.headB)); diff != "" {
				t.Errorf("getIntersectionNode()\n%s", diff)
			}
		})
	}
}

func Test_listLen(t *testing.T) {
	tests := []struct {
		name  string
		head  *common.ListNode
		wantI int
	}{
		{name: "empty"},
		{name: "[1,2,3,4,5]",
			head:  common.CreateList([]int{1, 2, 3, 4, 5}),
			wantI: 5,
		},
		{name: "[1]",
			head:  common.CreateList([]int{1}),
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := listLen(tt.head); gotI != tt.wantI {
				t.Errorf("listLen() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func Test_moveToNSteps(t *testing.T) {
	tests := []struct {
		name    string
		head    *common.ListNode
		steps   int
		wantRes *common.ListNode
	}{
		{name: "empty"},
		{name: "[1,2,3,4,5], 2",
			head:    common.CreateList([]int{1, 2, 3, 4, 5}),
			steps:   2,
			wantRes: common.CreateList([]int{3, 4, 5}),
		},
		{name: "[0,9,8,7,6,5,4,3,2,1], 3",
			head:    common.CreateList([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}),
			steps:   3,
			wantRes: common.CreateList([]int{7, 6, 5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantRes, moveToNSteps(tt.head, tt.steps)); diff != "" {
				t.Errorf("moveToNSteps()\n%s", diff)
			}
		})
	}
}

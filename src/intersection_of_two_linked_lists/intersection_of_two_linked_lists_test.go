package intersection_of_two_linked_lists

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	headA1, headB1, tail1 := list.CreateTwoIntersectedLists([]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3})
	headA2, headB2, tail2 := list.CreateTwoIntersectedLists([]int{1, 9, 1}, []int{3}, []int{2, 4})
	headA3, headB3, tail3 := list.CreateTwoIntersectedLists([]int{3}, []int{1, 9, 1}, []int{2, 4})
	tests := []struct {
		name  string
		headA *list.Node
		headB *list.Node
		want  *list.Node
	}{
		{name: "empty"},

		{name: "no intersections",
			headA: list.Create([]int{1, 2, 3}),
			headB: list.Create([]int{1, 2, 3})},

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

		{name: "A=[3] B=[1,9,1] tail=[2,4]",
			headA: headA3,
			headB: headB3,
			want:  tail3,
		},

		{name: "A=[2,6,4] B=[1,5]",
			headA: list.Create([]int{2, 6, 4}),
			headB: list.Create([]int{1, 5}),
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

func Test_moveToNSteps(t *testing.T) {
	tests := []struct {
		name    string
		head    *list.Node
		steps   int
		wantRes *list.Node
	}{
		{name: "empty"},
		{name: "[1,2,3,4,5], 2",
			head:    list.Create([]int{1, 2, 3, 4, 5}),
			steps:   2,
			wantRes: list.Create([]int{3, 4, 5}),
		},
		{name: "[0,9,8,7,6,5,4,3,2,1], 3",
			head:    list.Create([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}),
			steps:   3,
			wantRes: list.Create([]int{7, 6, 5, 4, 3, 2, 1}),
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

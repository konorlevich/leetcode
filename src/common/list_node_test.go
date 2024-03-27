package common

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func Test_CreateList(t *testing.T) {
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
			t.Parallel()
			if got := CreateList(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateListWithCycle(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		pos   int
		wantL *ListNode
	}{
		{name: "empty"},

		{name: "head = [3,2,0,-4], pos = 1",
			nums: []int{3, 2, 0, -4},
			pos:  1,
			wantL: func() *ListNode {
				cycled := &ListNode{Val: 2}
				cycled.Next = &ListNode{Val: 0,
					Next: &ListNode{Val: -4}}
				cycled.Next.Next.Next = cycled
				return &ListNode{
					Val:  3,
					Next: cycled,
				}
			}(),
		},

		{name: "head = [3,2,0,-4], pos = -1",
			nums:  []int{3, 2, 0, -4},
			pos:   -1,
			wantL: CreateList([]int{3, 2, 0, -4}),
		},

		{name: "head = [1,2], pos = 0",
			nums: []int{1, 2},
			pos:  0,
			wantL: func() *ListNode {
				cycled := &ListNode{Val: 1}
				cycled.Next = &ListNode{Val: 2, Next: cycled}
				return cycled
			}(),
		},

		{name: "head = [1,2], pos = -1",
			nums:  []int{1, 2},
			pos:   -1,
			wantL: CreateList([]int{1, 2}),
		},

		{name: "head = [1], pos = 0",
			nums: []int{1},
			pos:  0,
			wantL: func() *ListNode {
				cycled := &ListNode{Val: 1}
				cycled.Next = cycled
				return cycled
			}(),
		},

		{name: "head = [1], pos = -1",
			nums:  []int{1},
			pos:   -1,
			wantL: CreateList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantL, CreateListWithCycle(tt.nums, tt.pos)); diff != "" {
				t.Errorf("CreateListWithCycle()\n%s", diff)
			}
		})
	}
}

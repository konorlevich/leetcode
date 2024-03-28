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
			name: "[1,2,3]",
			nums: []int{1, 2, 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					}}}},
		{
			name: "[0,9,8,7,6,5,4,3,2,1]",
			nums: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: &ListNode{Val: 0,
				Next: &ListNode{Val: 9,
					Next: &ListNode{Val: 8,
						Next: &ListNode{Val: 7,
							Next: &ListNode{Val: 6,
								Next: &ListNode{Val: 5,
									Next: &ListNode{Val: 4,
										Next: &ListNode{Val: 3,
											Next: &ListNode{Val: 2,
												Next: &ListNode{Val: 1},
											}}}}}}}}}},
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

func Test_WrapNode(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		node *ListNode
		want *ListNode
	}{
		{name: "empty"},
		{
			name: "[1,2]+[3,4,5]", nums: []int{1, 2},
			node: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5},
						},
					}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := WrapNode(tt.nums, tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTwoIntersectedLists(t *testing.T) {

	tail1, tail2 := CreateList([]int{1, 2, 3}), CreateList([]int{3, 2, 1})
	tests := []struct {
		name string

		numsA    []int
		numsB    []int
		tailNums []int

		wantHeadA *ListNode
		wantHeadB *ListNode
		wantTail  *ListNode
	}{
		{name: "empty"},
		{name: "[4,5,6],[7,8,9,0], tail1",
			numsA:     []int{4, 5, 6},
			numsB:     []int{7, 8, 9, 0},
			tailNums:  []int{1, 2, 3},
			wantHeadA: WrapNode([]int{4, 5, 6}, tail1),
			wantHeadB: WrapNode([]int{7, 8, 9, 0}, tail1),
			wantTail:  tail1,
		},
		{name: "[7,8,9],[4,5,6,0], tail2",
			numsA:     []int{7, 8, 9},
			numsB:     []int{4, 5, 6, 0},
			tailNums:  []int{3, 2, 1},
			wantHeadA: WrapNode([]int{7, 8, 9}, tail2),
			wantHeadB: WrapNode([]int{4, 5, 6, 0}, tail2),
			wantTail:  tail2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeadA, gotHeadB, gotTail := CreateTwoIntersectedLists(tt.numsA, tt.numsB, tt.tailNums)
			if !reflect.DeepEqual(gotHeadA, tt.wantHeadA) {
				t.Errorf("CreateTwoIntersectedLists() gotHeadA = %v, want %v", gotHeadA, tt.wantHeadA)
			}
			if !reflect.DeepEqual(gotHeadB, tt.wantHeadB) {
				t.Errorf("CreateTwoIntersectedLists() gotHeadB = %v, want %v", gotHeadB, tt.wantHeadB)
			}
			if !reflect.DeepEqual(gotTail, tt.wantTail) {
				t.Errorf("CreateTwoIntersectedLists() gotTail = %v, want %v", gotTail, tt.wantTail)
			}
		})
	}
}

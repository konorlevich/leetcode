package list

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func Test_CreateList(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *Node
	}{
		{name: "empty"},
		{
			name: "[1,2,3]",
			nums: []int{1, 2, 3},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 3,
					}}}},
		{
			name: "[0,9,8,7,6,5,4,3,2,1]",
			nums: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: &Node{Val: 0,
				Next: &Node{Val: 9,
					Next: &Node{Val: 8,
						Next: &Node{Val: 7,
							Next: &Node{Val: 6,
								Next: &Node{Val: 5,
									Next: &Node{Val: 4,
										Next: &Node{Val: 3,
											Next: &Node{Val: 2,
												Next: &Node{Val: 1},
											}}}}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Create(tt.nums); !reflect.DeepEqual(got, tt.want) {
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
		wantL *Node
	}{
		{name: "empty"},

		{name: "head = [3,2,0,-4], pos = 1",
			nums: []int{3, 2, 0, -4},
			pos:  1,
			wantL: func() *Node {
				cycled := &Node{Val: 2}
				cycled.Next = &Node{Val: 0,
					Next: &Node{Val: -4}}
				cycled.Next.Next.Next = cycled
				return &Node{
					Val:  3,
					Next: cycled,
				}
			}(),
		},

		{name: "head = [3,2,0,-4], pos = -1",
			nums:  []int{3, 2, 0, -4},
			pos:   -1,
			wantL: Create([]int{3, 2, 0, -4}),
		},

		{name: "head = [1,2], pos = 0",
			nums: []int{1, 2},
			pos:  0,
			wantL: func() *Node {
				cycled := &Node{Val: 1}
				cycled.Next = &Node{Val: 2, Next: cycled}
				return cycled
			}(),
		},

		{name: "head = [1,2], pos = -1",
			nums:  []int{1, 2},
			pos:   -1,
			wantL: Create([]int{1, 2}),
		},

		{name: "head = [1], pos = 0",
			nums: []int{1},
			pos:  0,
			wantL: func() *Node {
				cycled := &Node{Val: 1}
				cycled.Next = cycled
				return cycled
			}(),
		},

		{name: "head = [1], pos = -1",
			nums:  []int{1},
			pos:   -1,
			wantL: Create([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantL, CreateWithCycle(tt.nums, tt.pos)); diff != "" {
				t.Errorf("CreateWithCycle()\n%s", diff)
			}
		})
	}
}

func Test_WrapNode(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		node *Node
		want *Node
	}{
		{name: "empty"},
		{
			name: "[1,2]+[3,4,5]", nums: []int{1, 2},
			node: &Node{
				Val: 3,
				Next: &Node{
					Val:  4,
					Next: &Node{Val: 5},
				},
			},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 3,
						Next: &Node{
							Val:  4,
							Next: &Node{Val: 5},
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

	tail1, tail2 := Create([]int{1, 2, 3}), Create([]int{3, 2, 1})
	tests := []struct {
		name string

		numsA    []int
		numsB    []int
		tailNums []int

		wantHeadA *Node
		wantHeadB *Node
		wantTail  *Node
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

func Test_listLen(t *testing.T) {
	tests := []struct {
		name  string
		head  *Node
		wantI int
	}{
		{name: "empty"},
		{name: "[1,2,3,4,5]",
			head:  Create([]int{1, 2, 3, 4, 5}),
			wantI: 5,
		},
		{name: "[1]",
			head:  Create([]int{1}),
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := Len(tt.head); gotI != tt.wantI {
				t.Errorf("listLen() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

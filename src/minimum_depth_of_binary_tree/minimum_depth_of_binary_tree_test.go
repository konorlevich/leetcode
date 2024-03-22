package minimum_depth_of_binary_tree

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
		want int
	}{
		{name: "empty"},

		{name: "3,9,20,null,null,15,7",
			root: &common.TreeNode{Val: 3,
				Left: &common.TreeNode{Val: 9},
				Right: &common.TreeNode{Val: 20,
					Left:  &common.TreeNode{Val: 15},
					Right: &common.TreeNode{Val: 7}}},
			want: 2,
		},
		{name: "2,null,3,null,4,null,5,null,6",
			root: &common.TreeNode{Val: 2,
				Right: &common.TreeNode{Val: 3,
					Right: &common.TreeNode{Val: 4,
						Right: &common.TreeNode{Val: 5,
							Right: &common.TreeNode{Val: 6}}}}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minNum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty"},
		{name: "1,2,3",
			nums: []int{1, 2, 3},
			want: 1},
		{name: "3,2,1",
			nums: []int{3, 2, 1},
			want: 1},
		{name: "3,2,-1",
			nums: []int{3, 2, -1},
			want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := minNum(tt.nums...); gotMin != tt.want {
				t.Errorf("minNum() = %v, want %v", gotMin, tt.want)
			}
		})
	}
}

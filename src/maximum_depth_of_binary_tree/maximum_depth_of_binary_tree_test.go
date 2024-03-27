package maximum_depth_of_binary_tree

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
		want int
	}{
		{name: "empty", want: 0},

		{name: "3,9,20,null,null,15,7",
			root: &common.TreeNode{
				Val:  3,
				Left: &common.TreeNode{Val: 9},
				Right: &common.TreeNode{Val: 20,
					Left:  &common.TreeNode{Val: 15},
					Right: &common.TreeNode{Val: 7}},
			},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

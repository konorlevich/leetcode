package balanced_binary_tree

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
		want bool
	}{
		{name: "empty", want: true},

		{name: "3,9,20,null,null,15,7",
			root: &common.TreeNode{Val: 3,
				Left: &common.TreeNode{Val: 9},
				Right: &common.TreeNode{Val: 20,
					Left:  &common.TreeNode{Val: 15},
					Right: &common.TreeNode{Val: 7}},
			},
			want: true,
		},

		{name: "1,2,2,3,3,null,null,4,4",
			root: &common.TreeNode{Val: 1,
				Left: &common.TreeNode{Val: 2,
					Left: &common.TreeNode{Val: 3,
						Left:  &common.TreeNode{Val: 4},
						Right: &common.TreeNode{Val: 4}},
					Right: &common.TreeNode{Val: 3}},
				Right: &common.TreeNode{Val: 2},
			},
			want: false,
		},
		//[1,2,2,3,null,null,3,4,null,null,4]
		{name: "1,2,2,3,null,null,3,4,null,null,4",
			root: &common.TreeNode{Val: 1,
				Left: &common.TreeNode{Val: 2,
					Left: &common.TreeNode{Val: 3,
						Left: &common.TreeNode{Val: 3}}},
				Right: &common.TreeNode{Val: 2,
					Right: &common.TreeNode{Val: 2,
						Right: &common.TreeNode{Val: 2}}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

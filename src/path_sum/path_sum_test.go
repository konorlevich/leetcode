package path_sum

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *common.TreeNode
		targetSum int
		want      bool
	}{
		{name: "empty", want: false},

		{name: "root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22",
			root: &common.TreeNode{Val: 5,
				Left: &common.TreeNode{Val: 4,
					Left: &common.TreeNode{Val: 11,
						Left:  &common.TreeNode{Val: 7},
						Right: &common.TreeNode{Val: 2}}},
				Right: &common.TreeNode{Val: 8,
					Left: &common.TreeNode{Val: 13},
					Right: &common.TreeNode{Val: 4,
						Right: &common.TreeNode{Val: 1}}}},
			targetSum: 22,
			want:      true},

		{name: "root = [1,2,3], targetSum = 5",
			root: &common.TreeNode{Val: 1,
				Left:  &common.TreeNode{Val: 2},
				Right: &common.TreeNode{Val: 3},
			},
			targetSum: 5,
			want:      false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := hasPathSum(tt.root, tt.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

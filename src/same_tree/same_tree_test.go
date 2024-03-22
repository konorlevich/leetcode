package same_tree

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *common.TreeNode
		q    *common.TreeNode
		want bool
	}{
		{name: "empty", want: true},

		{name: "p = [1,2,3], q = [1,2,3]",
			p:    &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2}, Right: &common.TreeNode{Val: 3}},
			q:    &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2}, Right: &common.TreeNode{Val: 3}},
			want: true,
		},

		{name: "p = [1,2], q = [1,null,2]",
			p:    &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2}},
			q:    &common.TreeNode{Val: 1, Right: &common.TreeNode{Val: 2}},
			want: false,
		},

		{name: "p = [1,2,1], q = [1,1,2]",
			p:    &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2}, Right: &common.TreeNode{Val: 1}},
			q:    &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 1}, Right: &common.TreeNode{Val: 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.p, tt.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

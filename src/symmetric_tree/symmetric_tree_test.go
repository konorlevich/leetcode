package symmetric_tree

import (
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
		want bool
	}{
		{name: "empty", want: true},

		{name: "1,2,2,3,4,4,3",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
					Left: &common.TreeNode{
						Val: 3},
					Right: &common.TreeNode{
						Val: 4}},
				Right: &common.TreeNode{
					Val: 2,
					Left: &common.TreeNode{
						Val: 4},
					Right: &common.TreeNode{
						Val: 3}},
			},
			want: true},

		{name: "1,2,2,null,3,null,3",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
					Right: &common.TreeNode{
						Val: 3}},
				Right: &common.TreeNode{
					Val: 2,
					Right: &common.TreeNode{
						Val: 3}},
			},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

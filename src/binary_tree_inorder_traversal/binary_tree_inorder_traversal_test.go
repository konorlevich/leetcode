package binary_tree_inorder_traversal

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

var (
	funcs = map[string]func(root *common.TreeNode) []int{
		"recursive": recursive,
		"iterative": iterative,
	}
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
		want []int
	}{
		{name: "empty"},
		{name: "1,null,2,3",
			root: &common.TreeNode{Val: 1, Right: &common.TreeNode{Val: 2, Left: &common.TreeNode{Val: 3}}},
			want: []int{1, 3, 2}},
		{name: "1", root: &common.TreeNode{Val: 1}, want: []int{1}},
		{name: "1,null,2",
			root: &common.TreeNode{Val: 1, Right: &common.TreeNode{Val: 2}},
			want: []int{1, 2}},
		{name: "2,3,null,1",
			root: &common.TreeNode{Val: 2, Left: &common.TreeNode{Val: 3, Left: &common.TreeNode{Val: 1}}},
			want: []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for fName, f := range funcs {
				t.Run(fName, func(t *testing.T) {
					if diff := cmp.Diff(tt.want, f(tt.root)); diff != "" {
						t.Errorf("inorderTraversal()\n%s", diff)
					}
				})
			}
		})
	}
}

func Benchmark_inorderTraversal(b *testing.B) {
	cases := map[string]*common.TreeNode{
		"1,null,2,3": {Val: 1, Right: &common.TreeNode{Val: 2, Left: &common.TreeNode{Val: 3}}},
		"2,3,null,1": {Val: 2, Left: &common.TreeNode{Val: 3, Left: &common.TreeNode{Val: 1}}},
	}
	for name, root := range cases {
		b.Run(name, func(b *testing.B) {
			for s, f := range funcs {
				b.Run(s, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						f(root)
					}
				})
			}
		})
	}
}

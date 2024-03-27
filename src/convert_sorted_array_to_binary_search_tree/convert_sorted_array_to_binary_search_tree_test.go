package convert_sorted_array_to_binary_search_tree

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		wantOneOf []*common.TreeNode
	}{
		{name: "empty", wantOneOf: []*common.TreeNode{nil}},

		{name: "-10,-3,0,5,9",
			nums: []int{-10, -3, 0, 5, 9},
			wantOneOf: []*common.TreeNode{
				{Val: 0,
					Left: &common.TreeNode{Val: -3,
						Left: &common.TreeNode{Val: -10},
					},
					Right: &common.TreeNode{Val: 9,
						Left: &common.TreeNode{Val: 5},
					},
				},
				{Val: 0,
					Left: &common.TreeNode{Val: -10,
						Right: &common.TreeNode{Val: -3},
					},
					Right: &common.TreeNode{Val: 5,
						Right: &common.TreeNode{Val: 9},
					},
				},
			}},

		{name: "1,3",
			nums: []int{1, 3},
			wantOneOf: []*common.TreeNode{
				{Val: 3,
					Left: &common.TreeNode{Val: 1},
				},
				{Val: 1,
					Right: &common.TreeNode{Val: 3},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sortedArrayToBST(tt.nums)
			diffs := make([]string, len(tt.wantOneOf))
			for i, node := range tt.wantOneOf {
				if diff := cmp.Diff(node, got); diff == "" {
					return
				} else {
					diffs[i] = diff
				}
			}
			for _, diff := range diffs {
				t.Errorf("sortedArrayToBST()\n%s", diff)
			}
		})
	}
}

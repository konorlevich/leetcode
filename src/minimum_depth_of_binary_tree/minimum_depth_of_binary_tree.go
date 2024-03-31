// Package minimum_depth_of_binary_tree
//
// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
package minimum_depth_of_binary_tree

import "github.com/konorlevich/leetcode/src/common"

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	lengths := make([]int, 0, 2)
	if root.Left != nil {
		lengths = append(lengths, minDepth(root.Left))
	}
	if root.Right != nil {
		lengths = append(lengths, minDepth(root.Right))
	}

	return 1 + minNum(lengths...)
}
func minNum(nums ...int) (min int) {
	for i, num := range nums {
		if i == 0 {
			min = num
			continue
		}
		if num < min {
			min = num
		}
	}
	return min
}

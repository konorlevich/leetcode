package maximum_depth_of_binary_tree

import "github.com/konorlevich/leetcode/src/common"

// maxDepth
//
// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	//return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

package balanced_binary_tree

import "github.com/konorlevich/leetcode/src/common"

// isBalanced
//
// # Given a binary tree, determine if it is height-balanced
//
// A height-balanced binary tree
// is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	leftDepth := maxDepth(root.Left)
	RightDepth := maxDepth(root.Right)
	depthDiff := 0
	if leftDepth > RightDepth {
		depthDiff = leftDepth - RightDepth
	} else {
		depthDiff = RightDepth - leftDepth
	}

	return depthDiff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}

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

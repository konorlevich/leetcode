// Package symmetric_tree
//
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
package symmetric_tree

import "github.com/konorlevich/leetcode/src/common"

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return root.Left == root.Right
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	return isSymmetric(&common.TreeNode{Val: root.Left.Val, Left: root.Left.Left, Right: root.Right.Right}) &&
		isSymmetric(&common.TreeNode{Val: root.Left.Val, Left: root.Right.Left, Right: root.Left.Right})
}

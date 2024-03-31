// Package same_tree
//
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
package same_tree

import "github.com/konorlevich/leetcode/src/common"

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if (p == nil) != (q == nil) {
		return false
	} else if p == nil {
		return true
	}
	if (p.Left == nil) != (q.Left == nil) {
		return false
	}
	if (p.Right == nil) != (q.Right == nil) {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

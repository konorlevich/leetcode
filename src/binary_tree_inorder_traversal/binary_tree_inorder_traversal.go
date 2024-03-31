// Package binary_tree_inorder_traversal
//
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// In in-order, we always recursively traverse the current node's left subtree;
// next, we visit the current node, and lastly, we recursively traverse the current node's right subtree.
package binary_tree_inorder_traversal

import "github.com/konorlevich/leetcode/src/common"

func recursive(root *common.TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	res = make([]int, 0, 3)

	if leftRes := recursive(root.Left); leftRes != nil {
		res = append(res, leftRes...)
	}
	res = append(res, root.Val)

	if rightRes := recursive(root.Right); rightRes != nil {
		res = append(res, rightRes...)
	}
	return res
}

func iterative(root *common.TreeNode) (res []int) {
	var stack []*common.TreeNode
	node := root
	fromStack := false
	for node != nil {
		if !fromStack {
			if node.Left != nil {
				stack = append(stack, node)
				node = node.Left
				fromStack = false
				continue
			}
		}
		res = append(res, node.Val)
		if node.Right != nil {
			node = node.Right
			fromStack = false
			continue
		}
		if len(stack) == 0 {
			node = nil
			continue
		}
		node = stack[len(stack)-1]
		if len(stack) == 1 {
			stack = []*common.TreeNode{}
		} else {
			stack = stack[:max(len(stack)-1, 0)]
		}
		fromStack = true
	}
	return res
}

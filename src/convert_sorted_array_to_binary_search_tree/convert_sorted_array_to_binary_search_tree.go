// Package convert_sorted_array_to_binary_search_tree
//
// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
//
// A height-balanced binary tree
// is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
package convert_sorted_array_to_binary_search_tree

import "github.com/konorlevich/leetcode/src/common"

func sortedArrayToBST(nums []int) (root *common.TreeNode) {
	numLen := len(nums)
	if numLen == 0 {
		return
	}
	if numLen == 1 {
		return &common.TreeNode{Val: nums[0]}
	}
	centerI := numLen / 2
	if numLen%2 == 0 {
		centerI--
	}
	leftNums := nums[:centerI]
	rightNums := nums[centerI+1:]
	root = &common.TreeNode{Val: nums[centerI],
		Left:  sortedArrayToBST(leftNums),
		Right: sortedArrayToBST(rightNums),
	}
	return
}

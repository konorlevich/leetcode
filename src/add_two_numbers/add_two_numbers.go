package add_two_numbers

import "github.com/konorlevich/leetcode/src/common"

// addTwoNumbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	newNode := &common.ListNode{Val: l1.Val + l2.Val}
	if newNode.Val >= 10 {
		newNode.Val -= 10
		newNode.Next = addTwoNumbers(&common.ListNode{Val: 1}, addTwoNumbers(l1.Next, l2.Next))
		return newNode
	}
	newNode.Next = addTwoNumbers(l1.Next, l2.Next)
	return newNode
}

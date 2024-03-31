// Package reverse_linked_list
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
package reverse_linked_list

import "github.com/konorlevich/leetcode/src/common"

// iterative
//
// As a first step we create a new node with current value.
// Then we wrap it iteratively with each Next.
func iterative(head *common.ListNode) (newHead *common.ListNode) {
	if head == nil {
		return nil
	}
	newHead = &common.ListNode{Val: head.Val}
	for head.Next != nil {
		newHead = &common.ListNode{
			Val:  head.Next.Val,
			Next: newHead}
		head = head.Next
	}
	return
}

// recursive
//
// See wrap
func recursive(head *common.ListNode) (newHead *common.ListNode) {
	return wrap(head, nil)
}

// wrap
//
// We recursively wrap current.Val as a next value for current.Next
func wrap(current, next *common.ListNode) (newNode *common.ListNode) {
	if current == nil {
		return next
	}
	return wrap(current.Next, &common.ListNode{
		Val:  current.Val,
		Next: next,
	})
}

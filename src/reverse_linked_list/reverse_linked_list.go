// Package reverse_linked_list
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
package reverse_linked_list

import (
	"github.com/konorlevich/leetcode/src/common/list"
)

// iterative
//
// As a first step we create a new node with current value.
// Then we wrap it iteratively with each Next.
func iterative(head *list.Node) (newHead *list.Node) {
	if head == nil {
		return nil
	}
	newHead = &list.Node{Val: head.Val}
	for head.Next != nil {
		newHead = &list.Node{
			Val:  head.Next.Val,
			Next: newHead}
		head = head.Next
	}
	return
}

// recursive
//
// See wrap
func recursive(head *list.Node) (newHead *list.Node) {
	return wrap(head, nil)
}

// wrap
//
// We recursively wrap current.Val as a next value for current.Next
func wrap(current, next *list.Node) (newNode *list.Node) {
	if current == nil {
		return next
	}
	return wrap(current.Next, &list.Node{
		Val:  current.Val,
		Next: next,
	})
}

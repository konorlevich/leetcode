// Package linked_list_cycle
//
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again
// by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
// Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
package linked_list_cycle

import "github.com/konorlevich/leetcode/src/common"

// hasCycleWithNodeBuffer
//
// We iterate over the list, checking if we know current node, save it otherwise.
//
// Straight-forward and memory-consuming solution
func hasCycleWithNodeBuffer(head *common.ListNode) bool {
	if head == nil {
		return false
	}
	nodes := []*common.ListNode{head}
	node := head.Next
	for node != nil {
		for _, savedNode := range nodes {
			if savedNode == node {
				return true
			}
		}
		nodes = append(nodes, node)
		node = node.Next
	}
	return false
}

// hasCycleWithPointers
//
// An approach with slow and fast pointers I came up with by myself.
// We create two pointers with `head` and `head.Next`, then we iterate over the list,
// slow pointer receives `head.Next`, fast one receives `head.Next.Next`.
// As soon as we see nil in slow or fast pointer, we return false.
// If slow == fast, return true.
func hasCycleWithPointers(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s, f := head, head.Next

	for s != nil && f != nil {
		if s == f {
			return true
		}
		if s.Next == nil || f.Next == nil {
			return false
		}
		s, f = s.Next, f.Next.Next
	}
	return false
}

// hasCycleWithPointers2
//
// Optimized and cleaned solution. We start with slow and fast pointer equal `head`.
// Each iteration `slow` steps to `.Next`, `fast` steps `.Next.Next`, then we check for nil or equality.
func hasCycleWithPointers2(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s, f := head, head

	for s != nil && f != nil && s.Next != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}

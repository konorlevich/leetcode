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

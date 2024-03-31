// Package intersection_of_two_linked_lists
//
// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
// If the two linked lists have no intersection at all, return null.
//
// The test cases are generated such that there are no cycles anywhere in the entire linked structure.
//
// Note that the linked lists must retain their original structure after the function returns.
//
// The inputs to the judge are given as follows (your program is not given these inputs):
//
// - intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
//
// - listA - The first linked list.
//
// - listB - The second linked list.
//
// - skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
//
// - skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
//
// The judge will then create the linked structure based on these inputs and pass the two heads,
// headA and headB to your program.
// If you correctly return the intersected node, then your solution will be accepted.
package intersection_of_two_linked_lists

import (
	"github.com/konorlevich/leetcode/src/common/list"
)

// getIntersectionNode
//
// At first, we have to equalize heads length. We measure its length and then move longer list forward.
// Then we iterate over both lists and check if nodes are equal
func getIntersectionNode(headA, headB *list.Node) *list.Node {
	hA, hB := headA, headB
	lenA, lenB := list.Len(hA), list.Len(hB)
	if lenA > lenB {
		hA = moveToNSteps(hA, lenA-lenB)
	} else if lenB > lenA {
		hB = moveToNSteps(hB, lenB-lenA)
	}

	for hA != nil && hB != nil {
		if hA == hB {
			return hA
		}
		hA, hB = hA.Next, hB.Next
	}
	return nil
}

// moveToNSteps returns nth node from the list
func moveToNSteps(head *list.Node, n int) (res *list.Node) {
	res = head
	for i := 0; i < n; i++ {
		res = res.Next
	}
	return
}

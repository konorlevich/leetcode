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

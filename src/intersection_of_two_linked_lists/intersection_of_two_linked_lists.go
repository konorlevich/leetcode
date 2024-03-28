package intersection_of_two_linked_lists

import "github.com/konorlevich/leetcode/src/common"

// getIntersectionNode
//
// At first, we have to equalize heads length. We measure its length and then move longer list forward.
// Then we iterate over both lists and check if nodes are equal
func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	hA, hB := headA, headB
	lenA, lenB := listLen(hA), listLen(hB)
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

// listLen counts list nodes
func listLen(head *common.ListNode) (i int) {
	p := head
	for p != nil {
		i++
		p = p.Next
	}
	return i
}

// moveToNSteps returns nth node from the list
func moveToNSteps(head *common.ListNode, n int) (res *common.ListNode) {
	res = head
	for i := 0; i < n; i++ {
		res = res.Next
	}
	return
}

package merge_two_sorted_lists

import (
	"github.com/konorlevich/leetcode/src/common/list"
)

// mergeTwoLists
//
// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
func mergeTwoLists(list1 *list.Node, list2 *list.Node) *list.Node {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list2.Next != nil {
		list1 = mergeTwoLists(list1, list2.Next)
	}
	n := list1
	for n.Next != nil && n.Next.Val <= list2.Val {
		n = n.Next
	}
	if list2.Val < n.Val {
		v := *n
		list1 = &list.Node{Val: list2.Val, Next: &v}
	} else {
		list2.Next = n.Next
		n.Next = list2
	}
	return list1
}

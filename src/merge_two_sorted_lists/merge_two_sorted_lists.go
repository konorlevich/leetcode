package merge_two_sorted_lists

import "github.com/konorlevich/leetcode/src/common"

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list2.Next != nil {
		list1 = mergeTwoLists(list1, list2.Next)
	}
	n := list1
	for ; n.Next != nil && n.Next.Val <= list2.Val; n = n.Next {

	}
	if list2.Val < n.Val {
		v := *n
		list1 = &common.ListNode{Val: list2.Val, Next: &v}
	} else {
		list2.Next = n.Next
		n.Next = list2
	}
	return list1
}

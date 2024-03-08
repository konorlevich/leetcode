package remove_duplicates_from_sorted_list

import "github.com/konorlevich/leetcode/src/common"

// deleteDuplicates
//
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
func deleteDuplicates(f func(head *common.ListNode) *common.ListNode, head *common.ListNode) *common.ListNode {
	return f(head)
}

func iterative(head *common.ListNode) *common.ListNode {
	n := head
	for n != nil && n.Next != nil {
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		}
		n = n.Next
	}
	return head
}

func recursive(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next != nil {
		head.Next = recursive(head.Next)
	}

	if head.Next.Val == head.Val {
		head.Next = head.Next.Next
	}
	return head
}

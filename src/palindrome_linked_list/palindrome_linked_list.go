// Package palindrome_linked_list
//
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
package palindrome_linked_list

import (
	"github.com/konorlevich/leetcode/src/common/list"
)

func withBuffer(head *list.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	listLen := list.Len(head)
	l := make([]*list.Node, 0, listLen)
	l = append(l, head)
	for head = head.Next; head != nil; head = head.Next {
		l = append(l, head)
	}
	i, ii := 0, listLen-1
	for i < ii {
		if l[i].Val != l[ii].Val {
			return false
		}
		i++
		ii--
	}
	return true
}

package common

// ListNode
//
// Sorted list
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList
//
// Creates a sorted list (ListNode), defined by the integer sequence
func CreateList(nums []int) (l *ListNode) {
	curNode := &ListNode{}
	for _, i := range nums {
		if l == nil {
			curNode.Val = i
			l = curNode
			continue
		}
		curNode.Next = &ListNode{Val: i}
		curNode = curNode.Next
	}
	return l
}

package common

type ListNode struct {
	Val  int
	Next *ListNode
}

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

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

// CreateListWithCycle
//
// # Creates a list (ListNode) with a cycle
//
// There is a cycle in a linked list
// if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, `pos` is used to denote the index of the node that tail's next pointer is connected to.
func CreateListWithCycle(nums []int, pos int) (l *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	curNode := &ListNode{}
	for _, num := range nums {
		if l == nil {
			curNode.Val = num
			l = curNode
			continue
		}
		curNode.Next = &ListNode{Val: num}
		curNode = curNode.Next
	}
	if pos < 0 {
		return l
	}
	ii := 0
	cycleNode := l
	for ii < pos {
		cycleNode = l.Next
		ii++
	}
	curNode.Next = cycleNode

	return l
}

package list

type Node struct {
	Val  int
	Next *Node
}

// Create
//
// Creates a sorted list (ListNode), defined by the integer sequence
func Create(nums []int) (l *Node) {
	return WrapNode(nums, l)
}

// CreateTwoIntersectedLists
//
// Creates two lists, with common tail
func CreateTwoIntersectedLists(numsA, numsB, tailNums []int) (*Node, *Node, *Node) {
	tail := WrapNode(tailNums, nil)
	return WrapNode(numsA, tail), WrapNode(numsB, tail), tail
}

func WrapNode(nums []int, node *Node) (head *Node) {
	head = node
	for i := len(nums) - 1; i >= 0; i-- {
		head = &Node{Val: nums[i],
			Next: head}
	}
	return
}

// CreateWithCycle
//
// # Creates a list (ListNode) with a cycle
//
// There is a cycle in a linked list
// if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, `pos` is used to denote the index of the node that tail's next pointer is connected to.
func CreateWithCycle(nums []int, pos int) (l *Node) {
	if len(nums) == 0 {
		return nil
	}
	curNode := &Node{}
	for _, num := range nums {
		if l == nil {
			curNode.Val = num
			l = curNode
			continue
		}
		curNode.Next = &Node{Val: num}
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

// Len counts list nodes
func Len(head *Node) (i int) {
	p := head
	for p != nil {
		i++
		p = p.Next
	}
	return i
}

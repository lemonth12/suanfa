package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPail(head *ListNode) bool {
	if head == nil {
		return false
	}
	ints := make([]int, 0)
	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}

	for i := 0; i < len(ints)/2; i++ {
		if ints[i] != ints[len(ints)-1-i] {
			return false
		}
	}
	return true

}

func main() {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	isPail(listNode)

}

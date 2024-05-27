package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var listNode1 *ListNode
	var listNode2 *ListNode
	listNode1.Val = 1
	listNode1.Next = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	listNode1.Val = 2
	listNode1.Next = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
}

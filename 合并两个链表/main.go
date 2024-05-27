package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	cur := &ListNode{}
	newList := cur
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			newList.Next = head2
			head2 = head2.Next
		} else {
			newList.Next = head1
			head1 = head1.Next
		}
		newList = newList.Next
	}
	if head1 != nil {
		newList.Next = head1
	} else {
		newList.Next = head2
	}

	return newList.Next

}

func main() {

	//listNode1.Val = 1
	listNode1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	listNode2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	node := merge(listNode1, listNode2)
	//fmt.Println(node)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}

}

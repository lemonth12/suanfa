package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(listNode *ListNode, k int) *ListNode {
	if k == 0 {
		return nil
	}
	ints := make([]int, 0)
	for listNode != nil {
		ints = append(ints, listNode.Val)
		listNode = listNode.Next
	}
	if len(ints) < k {
		return nil
	}
	newInts := ints[len(ints)-k:]
	newList := &ListNode{Val: ints[len(ints)-k], Next: nil}
	cur := newList

	for i := 1; i < len(newInts); i++ {
		cur.Next = &ListNode{Val: newInts[i]}
		cur = cur.Next
	}

	return newList
}

func main() {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	node := FindKthToTail(listNode, 0)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}

}

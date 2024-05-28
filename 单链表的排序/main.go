package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortInList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ints := make([]int, 0)
	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	newList := &ListNode{Val: ints[0], Next: nil}
	cur := newList
	for _, v := range ints[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return newList
}

func main() {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	node := sortInList(listNode)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}

}

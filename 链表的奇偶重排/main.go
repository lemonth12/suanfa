package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	// write code here

	ints := make([]int, 0)
	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}
	i := make([]int, 0)
	i2 := make([]int, 0)
	for k, v := range ints {
		if k%2 == 0 {
			i2 = append(i2, v)
		} else {
			i = append(i, v)
		}
	}
	i2 = append(i2, i...)
	var newlist *ListNode
	if len(i2) > 0 {
		newlist = &ListNode{Val: i2[0], Next: nil}
		cur := newlist

		for _, v := range i2[1:] {
			cur.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			cur = cur.Next
		}
	}
	return newlist
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
	node := oddEvenList(listNode)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

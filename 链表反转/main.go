package main

import "fmt"

type ListNode struct {
	Node int
	Next *ListNode
}

func reverseListNode(list *ListNode) *ListNode {
	newList := list
	var pre *ListNode
	for newList != nil {
		recodeNode := newList.Next
		newList.Next = pre
		pre = newList
		newList = recodeNode
	}
	return pre
}

func main() {
	var listNode ListNode
	listNode.Node = 1
	listNode.Next = &ListNode{
		Node: 2,
		Next: &ListNode{
			Node: 3,
			Next: nil,
		},
	}

	node := reverseListNode(&listNode)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}

}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	node := reverse(head1)
	node2 := reverse(head2)
	newHead := &ListNode{}
	cur := newHead

	j := 0
	for node != nil || node2 != nil {
		h := j
		if node != nil {
			//cur.Next = node
			//fmt.Println(node.Val)
			h += node.Val
			node = node.Next
		}
		if node2 != nil {
			//cur.Next = node2
			//fmt.Println(node2.Val)
			h += node2.Val
			node2 = node2.Next
		}
		j = 0
		if h > 9 {
			h -= 10
			j = 1
		}
		fmt.Println(h)
		cur.Next = &ListNode{Val: h}
		cur = cur.Next
	}
	if j > 0 {
		cur.Next = &ListNode{Val: j} // 添加最后一个进位节点
	}
	listNode := reverse(newHead.Next)
	return listNode
}

func reverse(head *ListNode) *ListNode {
	newHead := head
	var pre *ListNode
	var recodeNode *ListNode

	for newHead != nil {
		recodeNode = newHead.Next
		newHead.Next = pre
		pre = newHead
		newHead = recodeNode
	}
	return pre
}

func main() {

	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	listNode1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	node := addInList(listNode, listNode1)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}

}

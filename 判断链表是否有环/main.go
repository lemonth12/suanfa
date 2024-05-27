package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func JustCycle(head *ListNode) bool {
	cycleMap := make(map[*ListNode]int)

	for head != nil {
		if _, ok := cycleMap[head]; ok {
			return true
		}
		cycleMap[head] = 1
		head = head.Next

	}
	return false
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	cycle := JustCycle(node1)
	fmt.Println(cycle)
}

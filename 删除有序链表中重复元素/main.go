package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteRepeat(list *ListNode) *ListNode {
	if list == nil {
		return list
	}

	m := make(map[int]int)
	ints := make([]int, 0)
	for list != nil {
		if _, ok := m[list.Val]; !ok {
			m[list.Val] = 1
			ints = append(ints, list.Val)
		}
		list = list.Next
	}

	newList := &ListNode{Val: ints[0], Next: nil}
	cur := newList
	for _, v := range ints[1:] {
		cur.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
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
				Val:  2,
				Next: nil,
			},
		},
	}
	repeat := DeleteRepeat(listNode)
	for repeat != nil {
		fmt.Println(repeat)
		repeat = repeat.Next
	}

}

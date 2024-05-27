package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteRepeat(list *ListNode) *ListNode {
	if list == nil {
		return list
	}

	m := make(map[int]int)

	for list != nil {
		if _, ok := m[list.Val]; !ok {
			m[list.Val] = 1
		} else {
			m[list.Val] = 2
		}
		list = list.Next
	}

	ints := make([]int, 0)
	for k, v := range m {
		if v == 1 {
			ints = append(ints, k)
		}
	}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	var newList *ListNode
	if len(ints) != 0 {
		newList = &ListNode{Val: ints[0], Next: nil}
		cur := newList
		for _, v := range ints[1:] {
			cur.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			cur = cur.Next
		}
	}

	return newList
}

func main() {
	listNode := &ListNode{
		Val: 2,
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

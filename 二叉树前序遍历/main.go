package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func preorderTraversal(root *TreeNode) []int {
	ints := make([]int, 0)
	newpreorderTraversal(root, &ints)
	return ints
}

func newpreorderTraversal(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	newpreorderTraversal(root.Left, arr)
	newpreorderTraversal(root.Right, arr)
}

func main() {
	node := NewNode(1)
	node.Right = NewNode(2)
	node.Right.Left = NewNode(3)
	traversal := preorderTraversal(node)
	fmt.Println(traversal)
}

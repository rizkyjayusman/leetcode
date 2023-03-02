package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(preorderTraversal(&TreeNode{
		Val: 1,
	}))
}

func preorderTraversal(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}

	arr = append(arr, root.Val)
	arr = append(arr, preorderTraversal(root.Left)...)
	arr = append(arr, preorderTraversal(root.Right)...)
	return arr
}

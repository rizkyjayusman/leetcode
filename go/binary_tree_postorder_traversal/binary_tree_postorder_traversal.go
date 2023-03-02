package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(postorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))

	// fmt.Println(postorderTraversal(&TreeNode{}))
	// fmt.Println(postorderTraversal(&TreeNode{Val: 1}))
}

func postorderTraversal(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}

	arr = append(arr, postorderTraversal(root.Left)...)
	arr = append(arr, postorderTraversal(root.Right)...)
	arr = append(arr, root.Val)

	return arr
}

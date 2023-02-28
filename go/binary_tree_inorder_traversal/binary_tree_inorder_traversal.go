package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(inorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
	// fmt.Println(inorderTraversal(&TreeNode{}))
	// fmt.Println(inorderTraversal(&TreeNode{
	// 	Val: 1,
	// }))
}

func inorderTraversal(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}

	arr = append(arr, inorderTraversal(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)
	return arr
}

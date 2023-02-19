package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn := TreeNode{}
	tn.Val = 1

	tn1R := TreeNode{}
	tn1R.Val = 2

	tn1R2L := TreeNode{}
	tn1R.Left = &tn1R2L
	tn.Right = &tn1R
	fmt.Println(inorderTraversal(&tn))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, 1)
	return res
}

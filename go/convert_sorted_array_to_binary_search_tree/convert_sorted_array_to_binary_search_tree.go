package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	print(sortedArrayToBST([]int{1, 3}))
	// print(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func print(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	print(root.Left)
	print(root.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

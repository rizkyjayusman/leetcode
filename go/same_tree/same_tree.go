package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tnp := TreeNode{}
	tnp.Val = 1

	var tnq TreeNode

	fmt.Println(isSameTree(&tnp, &tnq))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return true
	} else {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	fmt.Printf("check (%v) (%v)\n", p.Val, q.Val)
	return left && right
}

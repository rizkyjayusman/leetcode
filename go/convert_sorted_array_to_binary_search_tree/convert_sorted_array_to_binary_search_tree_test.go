package main

import (
	"fmt"
	"testing"
)

var (
	csa ConvertSortedArrayToBinarySearchTree = ConvertSortedArrayToBinarySearchTree{}
)

func TestConvertSortedArrayToBinarySearchTree(t *testing.T) {
	expected := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
	}
	actual := csa.SortedArrayToBST([]int{1, 3})
	compare(expected, actual, t)
}

func TestConvertSortedArrayToBinarySearchTree2(t *testing.T) {
	expected := &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -10,
			Left: &TreeNode{
				Val: -3,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 9,
			},
		},
	}
	actual := csa.SortedArrayToBST([]int{-10, -3, 0, 5, 9})
	compare(expected, actual, t)
}

func compare(expected *TreeNode, actual *TreeNode, t *testing.T) {
	if expected == nil && actual == nil {
		return
	}

	if (expected == nil && actual != nil) || (expected != nil && actual == nil) {
		t.Error("FAIL ~ value is not match")
	}

	if expected.Val != expected.Val {
		t.Error("FAIL ~ value is not match")
	}

	fmt.Println(expected.Val, actual.Val)
	compare(expected.Left, actual.Left, t)
	compare(expected.Right, actual.Right, t)
}

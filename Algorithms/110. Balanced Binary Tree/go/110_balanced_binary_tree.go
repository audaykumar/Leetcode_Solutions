package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	res := lh - rh

	if res < 0 {
		res = res * -1
	}

	if res <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if lh < rh {
		return rh + 1
	} else {
		return lh + 1
	}
}

func main() {

	test1 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9, Left: nil, Right: nil},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
			Right: &TreeNode{7, nil, nil}},
	}

	fmt.Printf("Result: %v\n", isBalanced(test1))
}

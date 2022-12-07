package main

import (
	"github.com/benjaminliu/leetcode-go-2/structures"
	"math"
)

type TreeNode = structures.TreeNode

func main() {

}

func maxDepth(root *TreeNode) int {
	return helper(root)
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left)
	right := helper(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

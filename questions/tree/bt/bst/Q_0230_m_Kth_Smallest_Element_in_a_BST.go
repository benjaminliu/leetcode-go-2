package bst

import "github.com/benjaminliu/leetcode-go-2/structures"

type TreeNode = structures.TreeNode

var (
	idx    int
	result int
)

func kthSmallest(root *TreeNode, k int) int {
	idx = 0
	inOrderTraverse(root, k)
	return result
}

func inOrderTraverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if idx >= k {
		return
	}

	inOrderTraverse(root.Left, k)

	idx++
	if idx == k {
		result = root.Val
		return
	}

	inOrderTraverse(root.Right, k)
}

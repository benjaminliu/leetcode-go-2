package main

import (
	"github.com/benjaminliu/leetcode-go-2/structures"
)

type TreeNode = structures.TreeNode

var next *TreeNode

func flatten(root *TreeNode) {
	next = nil
	postOrderTraverse(root)
}

func postOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	postOrderTraverse(root.Right)
	postOrderTraverse(root.Left)

	root.Right = next
	root.Left = nil

	next = root
}

package main

import "strconv"

var myMap map[string]int
var myList []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	myMap = make(map[string]int)
	myList = make([]*TreeNode, 0)
	_ = postOrderTraverse(root)
	return myList
}

func postOrderTraverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := postOrderTraverse(root.Left)
	right := postOrderTraverse(root.Right)

	value := left + "-" + right + "-" + strconv.Itoa(root.Val)
	count := myMap[value]

	if count == 1 {
		myList = append(myList, root)
	}
	myMap[value] = count + 1
	return value
}

package main

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	helper(root)
	return maxDiameter
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left)
	right := helper(root.Right)

	diameter := left + right
	maxDiameter = max(maxDiameter, diameter)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

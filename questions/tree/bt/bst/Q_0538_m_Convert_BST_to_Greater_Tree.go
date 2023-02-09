package bst

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	rightRootLeft(root)
	return root
}

func rightRootLeft(root *TreeNode) {
	if root == nil {
		return
	}

	rightRootLeft(root.Right)

	sum += root.Val
	root.Val = sum

	rightRootLeft(root.Left)
}

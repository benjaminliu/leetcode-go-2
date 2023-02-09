package bst

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	if !isValid(root.Left, min, root) {
		return false
	}

	if !isValid(root.Right, root, max) {
		return false
	}

	return true
}

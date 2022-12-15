package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		idxMap[inorder[i]] = i
	}

	return helper(preorder, 0, len(preorder)-1, 0, idxMap)
}

func helper(preorder []int, preStart, preEnd, inStart int, idxMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}

	cur := preorder[preStart]
	curIdx := idxMap[cur]
	count := curIdx - inStart

	root := &TreeNode{Val: cur}
	root.Left = helper(preorder, preStart+1, preStart+count, inStart, idxMap)
	root.Right = helper(preorder, preStart+count+1, preEnd, curIdx+1, idxMap)
	return root
}

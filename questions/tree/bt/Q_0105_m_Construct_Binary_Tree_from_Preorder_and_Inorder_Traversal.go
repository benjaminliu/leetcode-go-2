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

	cur := preorder[preStart]
	root := &TreeNode{Val: cur}

	if preStart == preEnd {
		return root
	}

	idx := idxMap[cur]
	lastDescentIdx := idx - 1

	//idxDiff = count - 1
	idxDiff := lastDescentIdx - inStart

	root.Left = helper(preorder, preStart+1, preStart+1+idxDiff, inStart, idxMap)
	root.Right = helper(preorder, preStart+1+idxDiff+1, preEnd, idx+1, idxMap)
	return root
}

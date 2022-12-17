package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		idxMap[inorder[i]] = i
	}

	return helper(postorder, 0, len(postorder)-1, 0, idxMap)
}

func helper(postorder []int, postStart, postEnd, inStart int, idxMap map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	cur := postorder[postEnd]
	root := &TreeNode{Val: cur}

	if postStart == postEnd {
		return root
	}

	idx := idxMap[cur]
	lastDescentIdx := idx - 1

	//idxDiff = count - 1
	idxDiff := lastDescentIdx - inStart

	root.Left = helper(postorder, postStart, postStart+idxDiff, inStart, idxMap)
	root.Right = helper(postorder, postStart+idxDiff+1, postEnd-1, idx+1, idxMap)
	return root
}

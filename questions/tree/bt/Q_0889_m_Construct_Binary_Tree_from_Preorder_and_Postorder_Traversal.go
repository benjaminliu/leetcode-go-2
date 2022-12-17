package main

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i := 0; i < len(postorder); i++ {
		idxMap[postorder[i]] = i
	}

	return helper(preorder, 0, len(preorder)-1, 0, len(postorder)-1, idxMap)
}

func helper(preorder []int, preStart, preEnd int, postStart, postEnd int, idxMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	cur := preorder[preStart]
	root := &TreeNode{Val: cur}

	if preStart == preEnd {
		return root
	}

	leftChild := preorder[preStart+1]
	leftChildIdx := idxMap[leftChild]

	// leftChildIdxDiff = leftChildSize - 1
	leftChildIdxDiff := leftChildIdx - postStart

	root.Left = helper(preorder, preStart+1, preStart+1+leftChildIdxDiff,
		postStart, leftChildIdx, idxMap)
	root.Right = helper(preorder, preStart+1+leftChildIdxDiff+1, preEnd,
		leftChildIdx+1, postEnd-1, idxMap)
	return root
}

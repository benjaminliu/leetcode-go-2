package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{Val: nums[start]}
	}

	max, maxIdx := nums[start], start

	for i := start + 1; i <= end; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}

	root := &TreeNode{Val: max}
	root.Left = buildTree(nums, start, maxIdx-1)
	root.Right = buildTree(nums, maxIdx+1, end)
	return root
}

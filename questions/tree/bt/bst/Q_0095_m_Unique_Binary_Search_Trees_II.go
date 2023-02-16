package bst

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	res := []*TreeNode{}
	if lo > hi {
		res = append(res, nil)
		return res
	}

	if lo == hi {
		res = append(res, &TreeNode{Val: lo})
		return res
	}

	for i := lo; i <= hi; i++ {
		left := build(lo, i-1)
		right := build(i+1, hi)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				res = append(res, root)
			}
		}
	}

	return res
}

package bst

var memo []int

func numTrees(n int) int {
	memo = make([]int, n+1)
	return count(1, n)
}

func count(lo, hi int) int {
	if lo >= hi {
		return 1
	}

	key := hi - lo
	if memo[key] != 0 {
		return memo[key]
	}

	res := 0
	for i := lo; i <= hi; i++ {
		left := count(lo, i-1)
		right := count(i+1, hi)
		res += left * right
	}

	memo[key] = res
	return res
}

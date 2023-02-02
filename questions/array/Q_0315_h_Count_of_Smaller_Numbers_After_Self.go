package main

type Pair struct {
	Val int
	Idx int
}

var (
	temp  []Pair
	count []int
)

func countSmaller(nums []int) []int {
	n := len(nums)
	count = make([]int, n)
	temp = make([]Pair, n)

	arr := make([]Pair, n)

	for i := 0; i < n; i++ {
		arr[i] = Pair{Val: nums[i], Idx: i}
	}

	sort(arr, 0, n-1)

	return count
}

func sort(arr []Pair, lo, hi int) {
	if lo == hi {
		return
	}

	mid := lo + (hi-lo)/2
	sort(arr, lo, mid)
	sort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

func merge(arr []Pair, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = arr[i]
	}

	leftPartIdx, rightPartIdx := lo, mid+1

	for cur := lo; cur <= hi; cur++ {
		if leftPartIdx > mid {
			arr[cur] = temp[rightPartIdx]
			rightPartIdx++
		} else if rightPartIdx > hi {
			arr[cur] = temp[leftPartIdx]
			leftPartIdx++

			// right part start from mid +1 to hi
			// the number from mid + 1 to rightPartIdx (exclude) is smaller than leftPartIdx
			count[arr[cur].Idx] += rightPartIdx - (mid + 1)
		} else if temp[leftPartIdx].Val > temp[rightPartIdx].Val {
			arr[cur] = temp[rightPartIdx]
			rightPartIdx++
		} else {
			arr[cur] = temp[leftPartIdx]
			leftPartIdx++

			// right part start from mid +1 to hi
			// the number from mid + 1 to rightPartIdx (exclude) is smaller than leftPartIdx
			count[arr[cur].Idx] += rightPartIdx - (mid + 1)
		}
	}
}

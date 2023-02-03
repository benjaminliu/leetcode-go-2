package main

var (
	temp  []int
	count int
)

func reversePairs(nums []int) int {
	n := len(nums)
	temp = make([]int, n)
	count = 0
	sort(nums, 0, n-1)
	return count
}

func sort(nums []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	merge(nums, lo, mid, hi)
}

func merge(nums []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}

	rightPartIdx := mid + 1
	for i := lo; i <= mid; i++ {
		for rightPartIdx <= hi && nums[i] > nums[rightPartIdx]*2 {
			rightPartIdx++
		}

		count += rightPartIdx - (mid + 1)
	}

	leftPartIdx := lo
	rightPartIdx = mid + 1

	for cur := lo; cur <= hi; cur++ {
		if leftPartIdx > mid {
			nums[cur] = temp[rightPartIdx]
			rightPartIdx++
		} else if rightPartIdx > hi {
			nums[cur] = temp[leftPartIdx]
			leftPartIdx++
		} else if temp[leftPartIdx] > temp[rightPartIdx] {
			nums[cur] = temp[rightPartIdx]
			rightPartIdx++
		} else {
			nums[cur] = temp[leftPartIdx]
			leftPartIdx++
		}
	}

}

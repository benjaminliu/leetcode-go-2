package main

var tmp []int

func sortArray(nums []int) []int {
	tmp = make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2
	mergeSort(nums, lo, mid)
	mergeSort(nums, mid+1, hi)

	merge(nums, lo, mid, hi)
}

func merge(nums []int, lo, mid, hi int) {
	for i := lo; i <= mid; i++ {
		tmp[i] = nums[i]
	}

	left := lo
	right := mid + 1

	for i := lo; i <= hi; i++ {
		if left > mid {
			nums[i] = nums[right]
			right++
		} else if right > hi {
			nums[i] = tmp[left]
			left++
		} else if tmp[left] <= nums[right] {
			nums[i] = tmp[left]
			left++
		} else {
			nums[i] = nums[right]
			right++
		}
	}
}

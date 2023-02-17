package bst

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	lo := 0
	hi := len(nums) - 1

	k = len(nums) - k

	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}

	return -1
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]

	i := lo + 1
	j := hi

	for i <= j {
		for i < hi && nums[i] <= pivot {
			i++
		}

		for j > lo && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}

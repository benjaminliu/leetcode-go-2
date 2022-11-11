package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dupCount := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dupCount++
		} else {
			if dupCount > 0 {
				nums[i-dupCount] = nums[i]
			}
		}
	}

	return len(nums) - dupCount
}

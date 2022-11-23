package main

func main() {

}

func getModifiedArray(length int, updates [][]int) []int {
	nums := make([]int, length)

	for _, update := range updates {
		i := update[0]
		j := update[1]
		val := update[2]

		nums[i] += val
		if j+1 < len(nums) {
			nums[j+1] -= val
		}
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}

package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if target < sum {
			right = right - 1
		} else {
			left = left + 1
		}
	}

	return []int{0, 0}
}

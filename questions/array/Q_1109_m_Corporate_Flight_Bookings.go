package main

func main() {

}

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)

	for _, booking := range bookings {
		i := booking[0] - 1
		j := booking[1] - 1
		val := booking[2]

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

package main

func main() {

}

func carPooling(trips [][]int, capacity int) bool {
	maxStop := 0

	for _, trip := range trips {
		if maxStop < trip[2] {
			maxStop = trip[2]
		}
	}

	nums := make([]int, maxStop)

	for _, trip := range trips {
		val := trip[0]
		start := trip[1]
		end := trip[2]

		nums[start] += val
		if end < maxStop {
			nums[end] -= val
		}
	}

	if nums[0] > capacity {
		return false
	}

	for i := 1; i < maxStop; i++ {
		nums[i] += nums[i-1]
		if nums[i] > capacity {
			return false
		}
	}

	return true
}

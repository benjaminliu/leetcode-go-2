package main

func main() {

}

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{}
	res.nums = nums
	res.sums = make([]int, len(nums))
	res.sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		res.sums[i] = res.sums[i-1] + nums[i]
	}

	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == right {
		return this.nums[left]
	}

	return this.sums[right] - this.sums[left] + this.nums[left]
}

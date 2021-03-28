package main


func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res := nums[0], 0
	for p := 0; p < len(nums);p++ {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
	}
	return maxSum
}
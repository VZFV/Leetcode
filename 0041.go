package main

func firstMissingPositive(nums []int) int {
	// find missing number from 1, if cannot find, return that number
	numMap := make(map[int]bool)
	for i := 0; i <len(nums) ; i++ {
		numMap[nums[i]] = true
	}

	for i := 1; i <len(nums) + 1; i++{
		if _, find := numMap[i]; !find {
			return i
		}
	}
	return len(nums) + 1
}
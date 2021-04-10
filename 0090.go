package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	subsets,res := []int{}, [][]int{}
	sort.Ints(nums)
	generateSubsetsWithDup(nums, 0, subsets, &res)

	return res
}

func generateSubsetsWithDup (nums []int, index int, subsets []int, res *[][]int) { // use DFS
	copySubsets := make([]int, len(subsets)) // since subsets is always being change, so we need a copy of it
	copy(copySubsets, subsets)
	*res = append(*res, copySubsets)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		subsets = append(subsets, nums[i])
		generateSubsetsWithDup(nums, i + 1, subsets, res) // go to next level
		subsets = subsets[:len(subsets)-1]		// backtrack to previous level
	}
}
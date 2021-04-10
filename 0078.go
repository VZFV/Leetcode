package main

func subsets(nums []int) [][]int {
	subsets,res := []int{}, [][]int{}
	generateSubsets(nums, 0, subsets, &res)

	return res
}

func generateSubsets (nums []int, index int, subsets []int, res *[][]int) { // use DFS
	copySubsets := make([]int, len(subsets)) // since subsets is always being change, so we need a copy of it
	copy(copySubsets, subsets)
	*res = append(*res, copySubsets)
	for i := index; i < len(nums); i++ {
		subsets = append(subsets, nums[i])
		generateSubsets(nums, i + 1, subsets, res) // go to next level
		subsets = subsets[:len(subsets)-1]		// backtrack to previous level
	}
}
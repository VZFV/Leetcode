package main

import "fmt"

// time complexity = O(n), space complexity = O(n)
func twoSum(nums []int, target int) []int{
	m := make(map[int]int)					// initialize empty map
	// for each element in nums[], try to find target-nums[i] in map,
	// if find, return two indexes
	// if not, store the numbers and index into map until we found it
	for i := 0; i <len(nums);i++{
		another := target - nums[i]
		if _, isPresent := m[another]; isPresent {
			return []int{m[another],i}
		}
		m[nums[i]] = i
	}
	return nil
}


func main() {
	m := make(map[int]int)
	m[1] = 20
	val, isPresent := m[1]
	fmt.Println(val)
	fmt.Println(isPresent)
}

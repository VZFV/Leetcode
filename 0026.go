package main

import "fmt"

// time complexity: O(n), space complexity: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 1
	for i:=0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {	// if adjacent numbers are not equal, store the next numeber to current index position
			nums[index] = nums[i+1]
			index++
		}
	}
	return index
}

func main(){
	s := []int{0,1,2,2,3,4}
	result := removeDuplicates(s)
	fmt.Println(result)
	fmt.Println(s)
}
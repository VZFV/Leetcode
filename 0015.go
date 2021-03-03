package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(s)) // [[-1 -1 2] [-1 0 1]]

}


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0) 				// array to store result
	length := len(nums)
	low, high, sum := 0, 0, 0
	for index := 0; index < length-2; index ++ {
		if index == 0 || (index > 0 && nums[index] != nums[index-1]) {
			low = index + 1
			high = length - 1
			sum = 0 - nums[index]
			for low < high {
				if nums[low] + nums[high] == sum {
					result = append(result, []int{nums[index], nums[low], nums[high]})
					for low < high && nums[low] == nums[low+1] {  // ignore duplicate number
						low++
					}
					for low < high && nums[high] == nums[high-1] {  // ignore duplicate number
						high--
					}
					low++
					high--
				} else if nums[low] + nums[high] < sum{
					low++
				} else{
					high--
				}
			}
		}
	}
	return result
}


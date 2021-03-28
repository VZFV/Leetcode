package main

func searchRange(nums []int, target int) []int {
	if len(nums)==0 {
		return []int{-1,-1}
	}
	return []int{searchFirstElement(nums,target),searchLastElement(nums,target) }
}

func searchFirstElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		// find the mid point between low and high
		// x << y means x*2^y, x >> y means x/2^y
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid - 1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchLastElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		// find the mid point between low and high
		// x << y means x*2^y, x >> y means x/2^y
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums) - 1 || nums[mid + 1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

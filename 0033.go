package main

// use binary search to make sure time complexity O(logN)
func search(nums []int, target int) int {
	if len(nums)==0{
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		// find the mid point between low and high
		// x << y means x*2^y, x >> y means x/2^y
		mid := low + (high - low) >> 1
		if nums[mid] == target { // check if mid is target number
			return mid
		} else if nums[mid] > nums[low]{		// focus on larger side
			// if target is between low and high, move high to mid-1
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {	// else move low to mid + 1
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {		// focus on smaller side
			// if mid < target < high, move low to right side
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else { // else move high to left side
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
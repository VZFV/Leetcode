package main

func sortColors(nums []int)  {
	if len(nums) == 0 {
		return
	}
	red, white, blue := 0, 0, 0
	for _, value := range nums {
		if value == 0 {
			nums[blue] = 2
			blue++
			nums[white] = 1
			white++
			nums[red] = 0
			red++
		} else if value == 1 {
			nums[blue] = 2
			blue++
			nums[white] = 1
			white++
		} else {
			blue++
		}
	}
	//i, j, k := 0, 0, len(nums)-1
	//for j <= k {
	//	if nums[j] == 0 {
	//		nums[j], nums[i] = nums[i], nums[j]
	//		j++
	//		i++
	//	} else if nums[j] == 1 {
	//		j++
	//	} else if nums[j] == 2{
	//		nums[j], nums[k] = nums[k], nums[j]
	//		k--
	//	}
	//}

}

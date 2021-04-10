package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	// from end to save number
	i, j := m-1, n-1
	k := m + n - 1
	for ; k >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			i--
		} else {
			if nums1[i] < nums2[j] {
				nums1[k] = nums2[j]
				j--
			} else {
				nums1[k] = nums1[i]
				i--
			}
		}
	}
}
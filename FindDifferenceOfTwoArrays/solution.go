package finddifferenceoftwoarrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	vals1 := make(map[int]bool)
	vals2 := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		if _, present := vals1[nums1[i]]; !present {
			vals1[nums1[i]] = true
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, present := vals2[nums2[i]]; !present {
			vals2[nums2[i]] = true
		}
	}

	result := make([][]int, 2)
	result[0] = make([]int, 0)

	for val, _ := range vals1 {
		if _, present := vals2[val]; !present {
			result[0] = append(result[0], val)
		}
	}

	result[1] = make([]int, 0)

	for val, _ := range vals2 {
		if _, present := vals1[val]; !present {
			result[1] = append(result[1], val)
		}
	}

	return result
}

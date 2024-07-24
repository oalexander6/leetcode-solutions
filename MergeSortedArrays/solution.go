package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	if len(nums1) == 0 {
		copy(nums1, nums2)
		return
	}

	nums2I := 0

	for i := 0; i < len(nums1); i++ {
		num1 := nums1[i]
		num2 := nums2[nums2I]

		if num1 <= num2 {
			continue
		}

		copy(nums1[i+1:], nums1[i:len(nums1)-1])
		nums1[i] = num2
		nums2I++
	}

	// bumped := make([]int, 0)
	// nums2I := 0

	// for nums1I := 0; nums1I < len(nums1); nums1I++ {
	//     hasValidNum1 := nums1I + n > len(nums1)

	//     if hasValidNum1 && nums1[nums1I] >= num2 {
	//         continue
	//     }

	//     if hasValidNum1 && ()

	//     // if num1 > num2 and lastBumped > num2, bump num1, overwrite with num2, increment nums2I

	//     // if num1 > num2 and lastBumped < num2, bump num1, overwrite with lastBumped
	// }
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

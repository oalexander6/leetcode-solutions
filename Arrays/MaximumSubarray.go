package arrays

import "math"

func maxSubArray(nums []int) int {
	currMax := math.MinInt
	maxEndingHere := 0

	for i := 0; i < len(nums); i++ {
		maxEndingHere = maxEndingHere + nums[i]
		if currMax < maxEndingHere {
			currMax = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return currMax
}

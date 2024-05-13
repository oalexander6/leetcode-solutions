package arraypartition

import "slices"

// https://leetcode.com/problems/array-partition/submissions/1250451825

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	sum := 0

	for i := 0; i < len(nums)-1; i += 2 {
		sum += nums[i]
	}

	return sum
}

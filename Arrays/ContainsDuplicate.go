package arrays

func containsDuplicate(nums []int) bool {
	foundNums := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, found := foundNums[nums[i]]; found {
			return true
		}

		foundNums[nums[i]] = true
	}

	return false
}

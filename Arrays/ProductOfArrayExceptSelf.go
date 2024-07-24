package arrays

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1

	for i := 0; i < len(nums)-1; i++ {
		result[i+1] = result[i] * nums[i]
	}

	currProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = currProduct * result[i]
		currProduct = currProduct * nums[i]
	}

	return result
}

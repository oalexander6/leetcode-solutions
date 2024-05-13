package productofarrayexceptself

type entry struct {
	Prefix int
	Suffix int
}

func productExceptSelf(nums []int) []int {
	prefixProduct := 1
	suffixProduct := 1

	entries := make([]entry, len(nums))

	for offset := 0; offset < len(nums); offset++ {
		entries[offset].Prefix = prefixProduct
		entries[len(nums)-1-offset].Suffix = suffixProduct

		prefixProduct = prefixProduct * nums[offset]
		suffixProduct = suffixProduct * nums[len(nums)-1-offset]
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = entries[i].Prefix * entries[i].Suffix
	}

	return result
}

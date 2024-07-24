package arrays

func twoSum(nums []int, target int) []int {
	visitedNums := make(map[int]int)

	for i, num := range nums {
		pair := target - num

		if pairI, found := visitedNums[pair]; found {
			return []int{pairI, i}
		}

		visitedNums[num] = i
	}

	return []int{}
}

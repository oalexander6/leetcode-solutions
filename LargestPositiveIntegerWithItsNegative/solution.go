package largestpositiveintegerwithitsnegative

func findMaxK(nums []int) int {
	negatives := make(map[int]bool)
	positives := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		currNum := nums[i]

		if currNum < 0 {
			negatives[currNum] = true
		} else {
			positives[currNum] = true
		}
	}

	maxVal := -1

	for num, _ := range positives {
		if _, found := negatives[-1*num]; found && num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

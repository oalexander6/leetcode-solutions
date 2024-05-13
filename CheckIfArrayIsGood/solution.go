package checkifarrayisgood

func isGood(nums []int) bool {
	maxValue := 0
	numsFound := make(map[int]int)

	for _, val := range nums {
		if val > maxValue {
			maxValue = val
		}

		if _, alreadyFound := numsFound[val]; alreadyFound {
			numsFound[val] = numsFound[val] + 1
		} else {
			numsFound[val] = 1
		}
	}

	if len(nums) != maxValue+1 || numsFound[maxValue] != 2 {
		return false
	}

	for i := maxValue; i > 0; i-- {
		if _, found := numsFound[i]; !found {
			return false
		}
	}

	return true
}

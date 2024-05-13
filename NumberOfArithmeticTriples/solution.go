package numberofarithmetictriples

func arithmeticTriplets(nums []int, diff int) int {
	numTriplets := 0

	numsFound := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		currNum := nums[i]

		numsFound[currNum] = true

		prev1 := currNum - diff
		prev2 := prev1 - diff

		_, prev1Found := numsFound[prev1]
		_, prev2Found := numsFound[prev2]

		if prev1Found && prev2Found {
			numTriplets++
		}
	}

	return numTriplets
}

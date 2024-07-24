package main

func longestOnes(nums []int, k int) int {
	rightPtr := 0
	leftPtr := 0
	numFlipped := 0
	maxLength := 0

	for {
		if rightPtr >= len(nums) {
			break
		}

		if nums[rightPtr] == 1 {
			maxLength = max(maxLength, rightPtr-leftPtr+1)
			rightPtr++
			continue
		}

		if nums[rightPtr] == 0 && numFlipped < k {
			numFlipped++
			maxLength = max(maxLength, rightPtr-leftPtr+1)
			rightPtr++
			continue
		}

		// rightPtr is a 0 and we can not flip a new one
		found := false
		for {
			if leftPtr >= rightPtr {
				break
			}

			if nums[leftPtr] == 0 {
				leftPtr++
				found = true
				break
			}

			leftPtr++
		}

		if !found {
			rightPtr++
			leftPtr = rightPtr
		}
	}

	return maxLength
}

func main() {
	input := []int{0, 0, 1, 1, 1, 0, 0}

	_ = longestOnes(input, 0)
}

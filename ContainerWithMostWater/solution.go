package containerwithmostwater

func maxArea(height []int) int {
	leftPtr := 0
	rightPtr := len(height) - 1

	maxArea := 0

	for {
		if leftPtr >= rightPtr {
			break
		}

		width := rightPtr - leftPtr
		smallerHeight := min(height[leftPtr], height[rightPtr])
		area := width * smallerHeight

		maxArea = max(area, maxArea)

		if height[leftPtr] < height[rightPtr] {
			leftPtr++
		} else {
			rightPtr--
		}
	}

	return maxArea
}

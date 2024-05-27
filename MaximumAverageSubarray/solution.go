package maximumaveragesubarray

func findMaxAverage(nums []int, k int) float64 {
	currSum := 0

	for i := 0; i < k; i++ {
		currSum += nums[i]
	}
	maxAvg := float64(currSum) / float64(k)

	leftPtr := 0
	for rightPtr := k; rightPtr < len(nums); rightPtr++ {
		currSum -= nums[leftPtr]
		currSum += nums[rightPtr]

		avg := float64(currSum) / float64(k)

		if avg > maxAvg {
			maxAvg = avg
		}

		leftPtr++
	}

	return maxAvg
}

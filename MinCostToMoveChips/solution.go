package mincosttomovechips

func minCostToMoveChips(position []int) int {
	numEven := 0
	numOdd := 0

	for i := 0; i < len(position); i++ {
		if position[i]%2 == 0 {
			numEven++
		} else {
			numOdd++
		}
	}

	return min(numEven, numOdd)
}

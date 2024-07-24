package kidswithgreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	maxCandies := 0
	for _, numCandies := range candies {
		if numCandies > maxCandies {
			maxCandies = numCandies
		}
	}

	for i, numCandies := range candies {
		if numCandies+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}

package arrays

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		currProfit := prices[i] - minPrice

		if currProfit > maxProfit {
			maxProfit = currProfit
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}

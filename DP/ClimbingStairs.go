package dp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	oneBefore := 2
	twoBefore := 1
	curr := 0

	for i := 3; i <= n; i++ {
		curr = oneBefore + twoBefore

		twoBefore = oneBefore
		oneBefore = curr
	}

	return curr
}

package countoperationstoobtainzero

func countOperations(num1 int, num2 int) int {
	operations := 0

	for {
		if num1 == 0 || num2 == 0 {
			break
		}

		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		operations++
	}

	return operations
}

package binarynumberwithalternatingbits

func hasAlternatingBits(n int) bool {
	currN := n
	lastDigit := 1 & currN
	currN = currN >> 1

	for {
		if currN > -0 {
			if 1&currN == lastDigit {
				return false
			}
			lastDigit = 1 & currN
			currN = currN >> 1
		} else {
			break
		}
	}

	return true
}

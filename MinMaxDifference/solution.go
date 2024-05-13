package minmaxdifference

import (
	"strconv"
	"strings"
)

func minMaxDifference(num int) int {
	numString := strconv.Itoa(num)

	leadingDigit := string(numString[0])

	minVal, _ := strconv.Atoi(strings.Replace(numString, leadingDigit, "0", -1))

	firstNonNineIndex := -1

	for i := 0; i < len(numString); i++ {
		if string(numString[i]) != "9" {
			firstNonNineIndex = i
			break
		}
	}

	if firstNonNineIndex == -1 {
		return num - minVal
	}

	firstNonNineDigit := string(numString[firstNonNineIndex])

	maxVal, _ := strconv.Atoi(strings.Replace(numString, firstNonNineDigit, "9", -1))

	return maxVal - minVal
}

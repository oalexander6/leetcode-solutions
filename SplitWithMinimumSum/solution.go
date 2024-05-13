package splitwithminimumsum

import (
	"sort"
	"strconv"
	"strings"
)

func splitNum(num int) int {
	digits := strings.Split(strconv.Itoa(num), "")
	sort.Strings(digits)

	num1 := ""
	num2 := ""

	for i := 0; i < len(digits); i += 2 {
		num1 += digits[i]

		if i+1 < len(digits) {
			num2 += digits[i+1]
		}
	}

	num1Val, _ := strconv.Atoi(num1)
	num2Val, _ := strconv.Atoi(num2)

	return num1Val + num2Val
}

package arrayconcatenationvalue

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	leftIndex := 0
	rightIndex := len(nums) - 1

	result := int64(0)

	for {
		if leftIndex > rightIndex {
			break
		}

		if leftIndex == rightIndex {
			result += int64(nums[leftIndex])
			break
		}

		currConcatVal, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[leftIndex], nums[rightIndex]))
		result = result + int64(currConcatVal)

		leftIndex++
		rightIndex--
	}

	return result
}

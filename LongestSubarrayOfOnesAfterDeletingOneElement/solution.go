package main

func longestSubarray(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	firstOne := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			firstOne = i
			break
		}
	}

	if firstOne == -1 {
		return 0
	}

	right := firstOne + 1
	left := firstOne

	hasDeletedZero := false
	maxLength := 0

	for {
		if right >= len(nums) {
			break
		}

		if nums[right] == 1 {
			newLength := right - left
			if !hasDeletedZero {
				newLength++
			}
			maxLength = max(maxLength, newLength)
			right++
			continue
		}

		if !hasDeletedZero {
			hasDeletedZero = true
			newLength := right - left
			maxLength = max(maxLength, newLength)
			right++
			continue
		}

		for {
			if left >= right {
				break
			}

			if nums[left] == 0 {
				left++
				break
			}

			left++
		}
	}

	return maxLength
}

func main() {
	input := []int{1, 1, 1}

	_ = longestSubarray(input)
}

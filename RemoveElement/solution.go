package main

func removeElement(nums []int, val int) int {
	goodVals := 0

	for i, num := range nums {
		if num != val {
			goodVals++
			continue
		}

		swapWithIndex := -1

		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] != val {
				swapWithIndex = j
				break
			}
		}

		if swapWithIndex != -1 {
			nums[i] = nums[swapWithIndex]
			nums[swapWithIndex] = 0
			goodVals++
		}
	}

	return goodVals
}

func main() {
	_ = removeElement([]int{3, 2, 2, 3}, 3)
}

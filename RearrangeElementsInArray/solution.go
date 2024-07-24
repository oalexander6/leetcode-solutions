package rearrangeelementsinarray

func rearrangeArray(nums []int) []int {
	positives := make([]int, 0)
	negatives := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			negatives = append(negatives, nums[i])
		} else {
			positives = append(positives, nums[i])
		}
	}

	result := make([]int, 0)

	for i := 0; i < len(positives); i++ {
		result = append(result, positives[i])
		result = append(result, negatives[i])
	}

	return result
}

func rearrangeArray2(nums []int) []int {
	result := make([]int, len(nums))

	negPtr := 0
	posPtr := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			for {
				if nums[posPtr] > 0 {
					result[i] = nums[posPtr]
					posPtr++
					break
				}

				posPtr++
			}
		} else {
			for {
				if nums[negPtr] < 0 {
					result[i] = nums[negPtr]
					negPtr++
					break
				}

				negPtr++
			}
		}
	}

	return result
}

func haveOppositeSigns(num1 int, num2 int) bool {
	return (num1 < 0 && num2 > 0) || (num1 > 0 && num2 < 0)
}

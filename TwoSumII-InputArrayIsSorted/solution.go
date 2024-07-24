package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		currNum := numbers[i]
		lookingFor := target - currNum

		for searchI := i + 1; searchI < len(numbers); searchI++ {
			if numbers[searchI] == lookingFor {
				return []int{i + 1, searchI + 1}
			}
			if numbers[searchI] > lookingFor {
				break
			}
		}
	}

	return []int{}
}

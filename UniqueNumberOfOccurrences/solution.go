package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int, 0)

	for _, val := range arr {
		if _, present := counts[val]; !present {
			counts[val] = 1
		} else {
			counts[val] = counts[val] + 1
		}
	}

	countsFound := make(map[int]bool)

	for _, count := range counts {
		if _, present := countsFound[count]; present {
			return false
		} else {
			countsFound[count] = true
		}
	}

	return true
}

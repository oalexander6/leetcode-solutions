package oddstringdifference

import "slices"

// https://leetcode.com/problems/odd-string-difference/submissions/1241290730

func oddString(words []string) string {
	v1 := []int{}
	v1Count := 0
	v1Word := ""
	v2 := []int{}
	v2Count := 0
	v2Word := ""

	for _, word := range words {
		result := getDifferenceArray(word)

		if slices.Compare(result, v1) == 0 {
			v1Count++
			continue
		}

		if slices.Compare(result, v2) == 0 {
			v2Count++
			continue
		}

		if slices.Compare([]int{}, v1) == 0 {
			v1 = result
			v1Word = word
			v1Count++
			continue
		}

		v2 = result
		v2Word = word
		v2Count++
	}

	if v1Count == 1 {
		return v1Word
	}

	return v2Word
}

func getDifferenceArray(word string) []int {
	result := make([]int, len(word)-1)

	for i := 0; i < len(word)-1; i++ {
		result[i] = int(word[i+1] - word[i])
	}

	return result
}

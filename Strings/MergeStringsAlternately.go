package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	result := ""

	maxLength := max(len(word1), len(word2))

	for currChar := 0; currChar < maxLength; currChar++ {
		if currChar < len(word1) {
			result += string(word1[currChar])
		}

		if currChar < len(word2) {
			result += string(word2[currChar])
		}
	}

	return result
}

package longestcommonprefix

// https://leetcode.com/problems/longest-common-prefix/submissions/1227092296

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	result := ""
	referenceString := strs[0]

	for charIndex := 0; charIndex < len(referenceString); charIndex++ {
		stop := false

		for stringIndex := 1; stringIndex < len(strs); stringIndex++ {
			currString := strs[stringIndex]

			if charIndex > len(currString)-1 || currString[charIndex] != referenceString[charIndex] {
				stop = true
				break
			}
		}

		if stop {
			break
		}

		result += string(referenceString[charIndex])
	}

	return result
}

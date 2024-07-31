package reversewordsinastring

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	result := ""

	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]

		if i > 0 {
			result += " "
		}
	}

	return result
}

package keyboardrow

import "strings"

// https://leetcode.com/problems/keyboard-row/submissions/1250444941

func findWords(words []string) []string {
	charMap := map[string]uint8{
		"q": 1,
		"w": 1,
		"e": 1,
		"r": 1,
		"t": 1,
		"y": 1,
		"u": 1,
		"i": 1,
		"o": 1,
		"p": 1,
		"a": 2,
		"s": 2,
		"d": 2,
		"f": 2,
		"g": 2,
		"h": 2,
		"j": 2,
		"k": 2,
		"l": 2,
		"z": 3,
		"x": 3,
		"c": 3,
		"v": 3,
		"b": 3,
		"n": 3,
		"m": 3,
	}

	results := make([]string, 0)

	for wordIndex := 0; wordIndex < len(words); wordIndex++ {
		row := charMap[strings.ToLower(string(words[wordIndex][0]))]
		good := true
		for charIndex := 0; charIndex < len(words[wordIndex]); charIndex++ {
			if charMap[strings.ToLower(string(words[wordIndex][charIndex]))] != row {
				good = false
				break
			}
		}

		if good {
			results = append(results, words[wordIndex])
		}
	}

	return results
}

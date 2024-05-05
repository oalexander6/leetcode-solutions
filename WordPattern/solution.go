package wordpattern

import "strings"

// https://leetcode.com/problems/word-pattern/submissions/1246057604

func wordPattern(pattern string, s string) bool {
	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if i >= len(words) {
			return false
		}

		expectedWord, hasWord := charToWord[pattern[i]]
		expectedChar, hasChar := wordToChar[words[i]]

		if hasWord && words[i] != expectedWord {
			return false
		}

		if hasChar && pattern[i] != expectedChar {
			return false
		}

		if !hasWord {
			charToWord[pattern[i]] = words[i]
		}

		if !hasChar {
			wordToChar[words[i]] = pattern[i]
		}
	}

	return true
}

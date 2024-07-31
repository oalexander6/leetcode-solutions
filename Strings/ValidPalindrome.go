package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	leftPtr := 0
	rightPtr := len(s) - 1

	for {
		if leftPtr >= rightPtr {
			break
		}

		for {
			if leftPtr >= len(s) || isAlphaNumeric(rune(s[leftPtr])) {
				break
			}

			leftPtr++
		}

		for {
			if rightPtr < 0 || isAlphaNumeric(rune(s[rightPtr])) {
				break
			}

			rightPtr--
		}

		if leftPtr > rightPtr {
			break
		}

		if !strings.EqualFold(string(s[leftPtr]), string(s[rightPtr])) {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}

func main() {
	_ = isPalindrome(".,")
}

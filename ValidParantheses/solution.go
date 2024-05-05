package validparantheses

// https://leetcode.com/problems/valid-parentheses/submissions/1227096535

var closingMappings = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isValid(s string) bool {
	stack := make([]string, 0)

	for i := range s {
		currChar := string(s[i])
		if currChar == "(" || currChar == "{" || currChar == "[" {
			stack = append(stack, currChar)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != closingMappings[currChar] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

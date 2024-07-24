package main

func gcdOfStrings(str1 string, str2 string) string {
	commonPrefix := ""

	for i := 0; i < min(len(str1), len(str2)); i++ {
		if str1[i] == str2[i] {
			commonPrefix += string(str1[i])
		} else {
			break
		}
	}

	if len(commonPrefix) == 0 {
		return ""
	}

	for i := len(commonPrefix); i > 0; i-- {
		if doesDivide(str1, commonPrefix[0:i]) && doesDivide(str2, commonPrefix[0:i]) {
			return commonPrefix[0:i]
		}
	}

	return ""
}

func doesDivide(value string, divisor string) bool {
	if len(value)%len(divisor) != 0 {
		return false
	}
	for i := 0; i < len(value); i++ {
		divisorIndex := i % len(divisor)

		if value[i] != divisor[divisorIndex] {
			return false
		}
	}
	return true
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"

	_ = gcdOfStrings(str1, str2)
}

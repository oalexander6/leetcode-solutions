package strings

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	longest := ""

	for i := 0; i < len(s); i++ {
		longest = check(i, i, longest, s)

		longest = check(i, i+1, longest, s)
	}

	return longest
}

func check(l int, r int, currLongest string, s string) string {
	longest := currLongest
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if r-l+1 > len(currLongest) {
			longest = s[l : r+1]
		}
		l--
		r++
	}
	return longest
}

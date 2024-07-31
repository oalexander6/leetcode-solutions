package strings

func countSubstrings(s string) int {
	if len(s) == 1 {
		return 1
	}

	count := 0

	for i := 0; i < len(s); i++ {
		count += countPalindromes(i, i, s)
		count += countPalindromes(i, i+1, s)
	}

	return count
}

func countPalindromes(l int, r int, s string) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}

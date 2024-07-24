package issubsequence

func isSubsequence(s string, t string) bool {
	if len(s) == len(t) {
		return s == t
	}

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return len(s) == 0
	}

	sPtr := 0

	for tPtr := 0; tPtr < len(t); tPtr++ {
		if sPtr >= len(s) {
			break
		}
		if t[tPtr] == s[sPtr] {
			sPtr++
		}
	}

	return sPtr == len(s)
}

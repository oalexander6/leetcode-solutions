package maxvowels

func maxVowels(s string, k int) int {
    left := 0
    right := 0
    vowelCount := 0

    for right < k {
        if isVowel(string(s[right])) {
            vowelCount++
        }

        right++
    }

    maxVowels := vowelCount

    for right < len(s) {
        vowelDelta := 0

        if isVowel(string(s[right])) {
            vowelDelta++
        }
        if isVowel(string(s[left])) {
            vowelDelta--
        }

        vowelCount += vowelDelta

        maxVowels = max(maxVowels, vowelCount)

        right++
        left++
    }

    return maxVowels
}

func isVowel(s string) bool {
    return s == "a" || s == "e" || s == "i" || s == "o" || s == "u" 
}

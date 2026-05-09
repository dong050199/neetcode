func longestPalindrome(s string) string {
    // solve this using two pointer
	leftIdx := 0
	maxLen := 1
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l <= r {
			if l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			if r - l + 1 > maxLen {
				maxLen = r - l + 1
				leftIdx = l
			}
			l--
			r++
		}

		l , r = i, i + 1
		for l <= r {
			if l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			if r - l + 1 > maxLen {
				maxLen = r - l + 1
				leftIdx = l
			}
			l--
			r++
		}
	}

	return s[leftIdx:leftIdx+maxLen]
}

func longestPalindrome(s string) string {
	maxLeft := 0
	maxLen := 0

	for i := 0; i <= len(s); i++ {
		l, r := i, i
		for l <= r {
			if l < 0 || r >= len(s) || s[l] != s[r]  {
				break
			}
			if r - l + 1 > maxLen {
				maxLeft = l
				maxLen = r - l + 1
			}
			l--
			r++
		}

		l, r = i, i + 1
		for l <= r {
			if  l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			if r - l + 1 > maxLen {
				maxLeft = l
				maxLen = r - l + 1
			}
			l--
			r++
		}
	}
	return s[maxLeft:maxLeft + maxLen]
}

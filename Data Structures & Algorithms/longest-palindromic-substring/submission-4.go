func longestPalindrome(s string) string {
	resIdx := 0
	resLen := 1

	for i := 0; i < len(s); i++ {
		l, r := i, i

		for l > 0 && r < len(s) && s[r] == s[l] {
			if r - l + 1 > resLen {
				resLen = r - l + 1
				resIdx = l
			}
			l--
			r++
		}

		l, r = i, i + 1
		for l > 0 && r < len(s) && s[r] == s[l] {
			if r - l + 1 > resLen {
				resLen = r - l + 1
				resIdx = l
			}
			l--
			r++
		}
	}

	return s[resIdx:resIdx+resLen]
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)
	for l <= r {
		if s[l] - 'a' > 26 {
			l++
			continue
		}

		if s[r] - 'a' > 26 {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}
	return true
}

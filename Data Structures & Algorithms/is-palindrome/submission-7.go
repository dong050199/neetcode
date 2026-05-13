func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s) - 1
	for l <= r {
		if !isAlphanumeric(rune(s[r])) {
			r--
			continue
		}

		if !isAlphanumeric(rune(s[l])) {
			l++
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

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

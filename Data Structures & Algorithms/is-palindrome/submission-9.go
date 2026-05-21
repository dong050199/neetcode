func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s) - 1 

	for l <= r {
		if !isAlphanumeric(rune(s[l])) {
			l++
			continue
		}
		
		if !isAlphanumeric(rune(s[r])) {
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

func isAlphanumeric(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
} 

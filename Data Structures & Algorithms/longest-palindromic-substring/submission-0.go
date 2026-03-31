func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[1:]
		}
	}

	maxLen := 0
	var res string
	for i := 1 ; i < len(s) ; i ++ {
		l, r := i -1, i + 1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}

			mx := r - l + 1
			if mx > maxLen {
				res = s[l:r+1] 
				maxLen = mx
			}

			l--
			r++
		}

		l, r = i -1 , i
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}

			mx := r - l + 1
			if mx > maxLen {
				res = s[l:r+1] 
				maxLen = mx
			}

			l--
			r++
		}
	}

	return res
}



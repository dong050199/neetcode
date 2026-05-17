func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

    maxLeft := 0
	maxLength := 1

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l <= r {
			if l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			
			if (r - l + 1) > maxLength {
				maxLength = r - l + 1
				maxLeft = l
			} 

			l--
			r++
		}

		l , r = i, i+1
		for l <= r {
			if l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			
			if (r - l + 1) > maxLength {
				maxLength = r - l + 1
				maxLeft = l
			} 

			l--
			r++
		}
	}

	return s[maxLeft:maxLeft + maxLength]
}

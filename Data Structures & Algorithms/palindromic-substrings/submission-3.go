func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l <= r {
			if l < 0 || r >= len(s) || s[r] != s[l] {
				break
			}
			res++
			l--
			r++
		}

		l, r = i, i + 1 
		for l <= r {
			if l < 0 || r >= len(s) || s[r] != s[l] {
				break
			}
			res++
			l--
			r++
		}
	}

	return res
}

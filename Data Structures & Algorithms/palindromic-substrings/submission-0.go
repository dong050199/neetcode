func countSubstrings(s string) int {
	res := len(s)

	for i := 1; i < len(s); i ++ {
		l, r := i -1, i + 1

		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			res++
			l--
			r++
		}

		l, r = i-1, i

		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			res++
			l--
			r++
		}
	}
	return res
}

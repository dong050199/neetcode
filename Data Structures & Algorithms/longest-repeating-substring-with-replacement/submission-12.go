func characterReplacement(s string, k int) int {
	if len(s) == 1 {
		return 1
	}
	maxf := 0
	res := 1
	l := 0
	mp := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		mp[s[r]]++
		maxf = max(maxf, mp[s[r]])

		for (r - l + 1) - maxf > k {
			mp[s[l]]--
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}

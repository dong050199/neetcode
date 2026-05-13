func characterReplacement(s string, k int) int {
	l := 0
	maxF := 0
	mp := make(map[byte]int)
	res := 0
	for r := 0; r < len(s); r++ {
		mp[s[r]]++
		maxF = max(maxF, mp[s[r]])

		for (r - l + 1) > maxF + k {
			mp[s[l]]--
			l++
		}
		res = max(res, r - l + 1)
	}
	return res
}

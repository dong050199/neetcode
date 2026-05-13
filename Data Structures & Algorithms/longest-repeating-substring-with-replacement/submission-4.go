func characterReplacement(s string, k int) int {
	l := 0
	res := 0
	mp := make(map[byte]bool)
	for r := 0; r < len(s); r++ {
		mp[s[r]] = true
		for len(mp) > k {
			delete(mp, s[l])
			l++
		}
		res = max(res, r - l + 1)
	}
	return res
}

func characterReplacement(s string, k int) int {
	l := 0
	res := 0
	mp := make(map[byte]bool)
	for r := 0; r < len(s); r++ {
		mp[s[r]] = true
		if k == 0 {
			if len(mp) != 1 {
				res = max(res, r - l + 1)
			}
			continue
		}

		for len(mp) > k {
			delete(mp, s[l])
			l++
		}
		res = max(res, r - l + 1)
	}
	return res
}

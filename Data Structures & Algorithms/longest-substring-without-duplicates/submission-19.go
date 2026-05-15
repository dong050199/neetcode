func lengthOfLongestSubstring(s string) int {
	l := 0
	res := 0
	mp := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; !exist {
			mp[s[r]] = r
			res = max(res, r - l + 1)
			continue
		}

		for l <= mp[s[r]] {
			delete(mp, s[l])
			l++
		}
		mp[s[r]] = r
	}
	return res
}

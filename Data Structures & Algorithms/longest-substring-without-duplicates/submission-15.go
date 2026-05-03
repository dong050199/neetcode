func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	l := 0
	res := 0
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			mp[s[r]] = r
			for l <= mp[s[r]] {
				delete(mp, s[l])
				l++
			}
			mp[s[r]] = r
			continue
		}
		mp[s[r]] = r
		res = max(res, len(mp))
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	l := 0
	mp := make(map[byte]int)
	res := 0
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			for l < mp[s[r]] + 1 {
				delete(mp, s[l])
				l++
			}
		}
		res = max(res, r - l + 1)
		mp[s[r]] = r
	}
	return res
}

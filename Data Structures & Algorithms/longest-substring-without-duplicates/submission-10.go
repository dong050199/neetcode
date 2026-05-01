func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int) 
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			for i := l; l <= mp[s[i]]; i++ {
				delete(mp, s[l])
				l++
			}
			l = mp[s[r]] + 1
			mp[s[r]] = r
		}
		mp[s[r]] = r
		res = max(res, len(mp))
	}

	return res
}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int) 
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if idx, exist := mp[s[r]]; exist {
			for i := l; i <= idx; i++ {
				delete(mp, s[i])
			}
			l = idx + 1
		}
		mp[s[r]] = r 
		res = max(res, len(mp))
	}

	return res
}
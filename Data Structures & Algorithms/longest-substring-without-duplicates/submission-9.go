func lengthOfLongestSubstring(s string) int {
	// we can use only one hash map
	mp := make(map[byte]int) 
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			for i := l; l <= mp[s[i]]; l++ {
				delete(mp, s[l])
			}
			l = mp[s[r]] + 1
			mp[s[r]] = r
		}
		mp[s[r]] = r
		res = max(res, len(mp))
	}

	return res
}

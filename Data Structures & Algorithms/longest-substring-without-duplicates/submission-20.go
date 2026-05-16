func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	res := 1

	l := 0 
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

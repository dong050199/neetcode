func lengthOfLongestSubstring(s string) int {
	res := 0
	for i := 0; i < len(s); i ++ {
		mp := make(map[byte]bool)
		mp[s[i]] = true
		count := 1
		for j := i + 1; j < len(s); j ++ {
			if !mp[s[j]] {
				mp[s[j]] = true
				count++
				continue
			}
			break
		}
		res = max(res, count)
	}

	return res
}

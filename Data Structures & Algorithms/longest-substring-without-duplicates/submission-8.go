func lengthOfLongestSubstring(s string) int {
	// brute force approach
	res := 0
	for i := 0; i < len(s); i++ {
		mp := make(map[byte]int)
		counter := 0
		for j := i; j < len(s); j++ {
			mp[s[j]]++
			if mp[s[j]] > 1 {
				break
			}
			counter++
			res = max(res, counter)
		}
	}
	return res
}

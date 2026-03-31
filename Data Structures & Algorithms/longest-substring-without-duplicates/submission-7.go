func lengthOfLongestSubstring(s string) int {
	left , right := 0, 0

	mp := make(map[byte]int)

	res := 0

	for right = 0; right < len(s); right ++ {
		if i, exist := mp[s[right]]; exist {
			left = max(left, i + 1)
		}
		mp[s[right]] = right
		res = max(res, right - left + 1)
	}
	return res
}

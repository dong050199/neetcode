func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
    // brute force approach
	minLen := 201
	for _, s := range strs {
		minLen = min(minLen, len(s))
	}

	res := 0
	for i := 0; i < minLen; i++ {
		cur := strs[0][i]
		for _, s := range strs {
			if s[i] != cur {
				return strs[0][:res]
			}
		}
		res++
	}

	return strs[0][:res]
}

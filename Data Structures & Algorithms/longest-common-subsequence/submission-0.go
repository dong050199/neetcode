func longestCommonSubsequence(text1 string, text2 string) int {
	// brute force approach, using recursive 

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if text1[i] == text2[j] {
			return 1 + dfs(i +1, j + 1)
		}

		return max(dfs(i + 1, j), dfs(i, j+1))
	}



	return dfs(0, 0)
}

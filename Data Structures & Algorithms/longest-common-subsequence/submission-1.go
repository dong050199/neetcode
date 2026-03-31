func longestCommonSubsequence(text1 string, text2 string) int {
	// brute force using recursive but we cache the result to reuse.

	memo := make([][]int, len(text1))

	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[i] {
			memo[i][j] = -1 
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if text1[i] == text2[j] {
			return 1 + dfs(i+1, j + 1)
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		memo[i][j] = max(dfs(i , j + 1), dfs(i + 1, j))

		return memo[i][j]
	}

	return dfs(0, 0)
}

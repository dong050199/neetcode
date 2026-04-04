func minDistance(word1 string, word2 string) int {
	var dfs func(i, j int) int 
	dfs = func(i, j int) int {
		if i == m {
			return n - j
		}

		if j == n {
			return m - i
		}

		if word1[i] == word2[j] {
			return dfs(i+1 , j+1)
		}

		res := min(dfs(i + 1, j) , dfs(i, j+1), dfs(i + 1, j +1))

	}


	return dfs(0, 0)
}

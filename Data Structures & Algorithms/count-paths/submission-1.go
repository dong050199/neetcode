func uniquePaths(m int, n int) int {
    // brute force approach, recursive
	var dfs func(i, j int) int
    dfs = func(i, j int) int {
		if i == m - 1 && j == n - 1 {
			return 1
		}

		if i >= m || j >= n {
			return 0
		}
		return dfs(i, j + 1) + dfs(i + 1, j)
	}

	return dfs(0, 0)
}

func numDistinct(s string, t string) int {
	if len(t) > len(s) {
		return 0
	}

	memo := make(map[[2]int]int)
	var dfs func (i, j int) int 
	dfs = func (i, j int) int {
		memoKey := [2]int{i, j}

		if _, exist := memo[memoKey] ; exist{
			return memo[memoKey]
		}

		if j == len(t) {
			return 1
		}

		if i == len(s) {
			return 0
		}

		if s[i] != t[j] {
			memo[memoKey] = dfs(i + 1, j)
			return memo[memoKey] 
		}

		memo[memoKey] = dfs(i + 1, j + 1) + dfs(i + 1, j)
		return memo[memoKey]
	}

	return dfs(0, 0)
}

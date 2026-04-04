func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	var res int
	dfs(n , 0, &res)
	return res
}

func dfs(n int, cur int, res *int) {
	if cur > n {
		return
	}

	if cur == n {
		*res = *res + 1
		return
	}

	dfs(n , cur + 1, res)
	dfs(n , cur + 2, res)
}

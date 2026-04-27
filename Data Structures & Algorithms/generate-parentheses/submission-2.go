func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(i, j int, cur string)

	dfs = func(i, j int, cur string) {
		if i == j && i == n {
			res = append(res, cur)
			return
		}

		if i < n {
			dfs(i + 1, j, cur + "(")
		}

		if j < i {
			dfs(i, j+1, cur + ")")
		}
	}

	dfs(0, 0, "")


	return res
}

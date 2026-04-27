func generateParenthesis(n int) []string {
	res := []string{} 
	
	var dfs func(open, close int, cur string)
	dfs = func(open, close int, cur string) {
		if open == n && close == n {
			res = append(res, cur)
			return
		}
		
		if open < n {
			dfs(open+1, close, cur+"(")
		}
		
		if close < open {
			dfs(open, close+1, cur+")")
		}
	}

	dfs(0, 0, "")

	return res
}
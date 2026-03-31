func generateParenthesis(n int) []string {
	var res []string
	dfs(n, 0, 0 ,"", &res)
	return res
}

func dfs(n int, open int, close int, cur string, res *[]string) {
    if len(cur) == 2*n {
        *res = append(*res, cur)
        return
    }

    if open < n {
        dfs(n, open+1, close, cur+"(", res)
    }

    if close < open {
        dfs(n, open, close+1, cur+")", res)
    }
}
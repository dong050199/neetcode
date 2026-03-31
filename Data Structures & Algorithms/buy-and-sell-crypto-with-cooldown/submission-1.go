func maxProfit(prices []int) int {
	var dfs func (i int, buying bool) int 
	dfs = func (i int, buying bool) int {
		if i >= len(prices) {
			return 0
		}

		cooldonw := dfs(i + 1, buying)
		if buying {
			buy := dfs(i + 1, false) - prices[i]
			return max(cooldonw, buy)
		} else {
			sell := dfs(i + 2, true) + prices[i]
			return max(cooldonw, sell) 
		}
	}

	return dfs(0, true)
}

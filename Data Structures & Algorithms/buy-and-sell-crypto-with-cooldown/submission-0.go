func maxProfit(prices []int) int {
	var dfs func(i int, buying bool) int
	dfs = func(i int, buying bool) int {
		if i >= len(prices) {
			return 0
		}

		cooldown := dfs(i+1, buying)
		if buying {
			buy := dfs(i + 1, false) - prices[i]
			return max(buy, cooldown)
		}else {
			sell := dfs(i + 2, true) + prices[i]
			return max(sell, cooldown)
		}
	}

    return dfs(0, true)
}

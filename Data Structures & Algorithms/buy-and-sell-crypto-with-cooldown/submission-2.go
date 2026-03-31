func maxProfit(prices []int) int {
	dp := make(map[[2]int]int)
	var dfs func(i int, canBuy bool) int 
	dfs = func(i int, canBuy bool) int {
		if i >= len(prices) {
			return 0
		}

		var dpKey [2]int
		if canBuy {
			dpKey = [2]int{i, 1}
		} else {
			dpKey = [2]int{i, 0}
		}

		if _, exist := dp[dpKey]; exist {
			return dp[dpKey]
		}

		cooldown := dfs(i + 1, canBuy)
		if canBuy {
			buy := dfs(i + 1, false) - prices[i]
			m := max(buy, cooldown)
			dp[[2]int{i, 1}] = m
			return m
		} else {
			sell := dfs(i + 2, true) + prices[i]
			m := max(sell, cooldown)
			dp[[2]int{i, 0}] = m
			return m
		}
	}
	return dfs(0, true)
}
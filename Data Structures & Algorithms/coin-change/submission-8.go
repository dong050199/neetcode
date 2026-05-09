func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i - c >= 0 {
				dp[i] = max(dp[i], dp[i-c] + 1)
			}
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}

	return dp[amount]
}

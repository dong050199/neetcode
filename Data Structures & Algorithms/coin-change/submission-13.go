func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)

	for a := 0; a <= amount; a++ {
		dp[a] = amount + 1
	}

	for a := 0; a <= amount; a++ {
		for _, c := range coins {
			if a - c >= 0 {
				dp[a] = min(dp[a], dp[a-c])
			}
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}

	return dp[amount]
}

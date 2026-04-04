func climbStairs(n int) int {
	dp := make([]int, n + 1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
	}

	return dp[n]
}

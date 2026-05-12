func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		one, _ := strconv.Atoi(s[i-2:i-1])
		two, _ := strconv.Atoi(s[i-2:i])
		if one > 0 {
			dp[i] += dp[i-1]
		}

		if two >= 10 && two <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

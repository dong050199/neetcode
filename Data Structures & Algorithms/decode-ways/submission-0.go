func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s) + 1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		oneDigit, _ := strconv.Atoi(s[i-1:i])
		twoDigit, _ := strconv.Atoi(s[i-2:i])

		if oneDigit != 0 { 
			dp[i] += dp[i-1]
		}

		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
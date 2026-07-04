func tribonacci(n int) int {
	// recursion not work as time limit, we need to make it more efficent.
	// switch n {
	// case 0:
	// 	return 0
	// case 1:
	// 	return 1
	// case 2:
	// 	return 1
	// default:
	// 	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	// }

	// dynamic programing
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
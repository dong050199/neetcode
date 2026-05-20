func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums) + 1)

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			} 

			res = max(res, dp[i])
		}
	}

	return res
}

func maxSubArray(nums []int) int {
	var dfs func(i int, flag bool) int

	dfs = func(i int, flag bool) int {
		if i >= len(nums) {
			if flag == true {
				return 0
			}
			return -10001
		}

		if flag {
			return max(0, dfs(i+1, true) + nums[i])
		}
		return max(dfs(i+1, false), nums[i] + dfs(i+1, true))
	}
	return dfs(0, false)
}

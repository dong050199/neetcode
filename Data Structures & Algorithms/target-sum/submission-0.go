func findTargetSumWays(nums []int, target int) int {
	var dfs func(i, am int) int 

	dfs = func(i, am int) int {
		if i == len(nums) {
			if am == target {
				return 1
			}
			return 0
		}

		return dfs(i + 1, am + nums[i]) + dfs(i+1, am - nums[i])
	}

	return dfs(0, target)
}

func findTargetSumWays(nums []int, target int) int {
	var dfs func(i, am int) int 
	memo := make(map[[2]int]int)
	dfs = func(i, am int) int {
		memoKey := [2]int{i, am}

		if _, exist := memo[memoKey]; exist {
			return memo[memoKey]
		}

		if i == len(nums) {
			if am == target {
				return 1
			}
			return 0
		}

		memo[memoKey] = dfs(i + 1, am + nums[i]) + dfs(i+1, am - nums[i])
		return memo[memoKey]
	}

	return dfs(0, 0)
}

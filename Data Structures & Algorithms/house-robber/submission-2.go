func rob(nums []int) int {
    return dfs(nums, 0)
}

func dfs(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	return max(dfs(nums, i+1), nums[i] + dfs(nums, i+2))
}
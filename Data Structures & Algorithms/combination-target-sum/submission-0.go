func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	dfs(nums, 0, target, []int{}, &res)
	return res
}

func dfs(nums []int, i int, target int, subset []int, res *[][]int) {
	total := 0
	for _, s := range subset {
		total += s
	}

	if target == total {
		tmp := make([]int, len(subset))
		copy(tmp, subset)
		*res = append(*res, tmp)
		return
	}

	if i >= len(nums) || total > target {
		return
	}

	subset = append(subset, nums[i])
	dfs(nums, i, target, subset, res)

	subset = subset[:len(subset)-1]
	dfs(nums, i+1, target, subset, res)
}
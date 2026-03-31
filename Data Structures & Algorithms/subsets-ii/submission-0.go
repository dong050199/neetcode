func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	dfs(nums, []int{}, 0, &res)
	return res
}

func dfs(nums []int, subset []int, i int, res *[][]int) {
	if i >= len(nums) {
		tmp := append([]int{}, subset... )
		*res = append(*res, tmp)
		return
	}

	subset = append(subset, nums[i])
	dfs(nums, subset, i+1, res)

	

	subset = subset[:len(subset) - 1]

	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}

	dfs(nums, subset, i+1, res)
}
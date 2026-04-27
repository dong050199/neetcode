func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, cur []int)
	dfs = func(i int, cur []int) {
		if i >= len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return 
		}

		cur = append(cur, nums[i])
		dfs(i + 1, cur)
		cur = cur[:len(cur)-1]
		dfs(i + 1, cur)
	}		

	dfs(0, []int{})
	return res
}

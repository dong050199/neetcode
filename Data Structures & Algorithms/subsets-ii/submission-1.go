func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var dfs func(i int, cur []int)

	dfs = func(i int, cur []int) {
		res = append(res, append([]int{}, cur...))

		for j := i ; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}

			cur = append(cur, nums[j])
			dfs(j + 1, cur)
			cur = cur[:len(cur) - 1]
		}
	}

	dfs(0, []int{})
	return res
}

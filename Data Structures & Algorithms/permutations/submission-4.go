func permute(nums []int) [][]int {
	res := [][]int{}
	selected := make([]bool, len(nums))
	var dfs func(cur []int)
	dfs = func(cur []int) {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		} 

		for i := 0; i < len(nums); i ++ {
			if selected[i] {
				continue
			}

			cur = append(cur, nums[i])
			selected[i] = true
			dfs(cur)
			cur = cur[:len(cur)-1]
			selected[i] = false
		}
	}

	dfs([]int{})

	return res
}

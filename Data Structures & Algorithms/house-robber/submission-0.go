func rob(nums []int) int {
	var res int

	dfs(nums, 0, make(map[int]bool), &res)

	return res
}


func dfs(nums []int, cur int, mp map[int]bool, res *int ) {
	if len(mp) == len(nums) {
		*res = max(*res, cur)
		return
	}

	for i, n := range nums {
		if !mp[i] {
			mp[i] = true
			if i >= 1 {
				mp[i-1] = true
			}
			if i < len(nums) - 1 {
				mp[i+1] = true
			}
			cur = cur + n
			dfs(nums, cur, mp, res)
		}
	}
}
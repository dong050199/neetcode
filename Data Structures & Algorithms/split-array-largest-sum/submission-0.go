func splitArray(nums []int, k int) int {
	var dfs func(i, arrLeft int) int
	cache := make(map[[2]int]int)
	dfs = func(i, arrLeft int) int {
		if i == len(nums) {
			if arrLeft == 0 {
				return 0
			}
			return math.MaxInt64
		}

		if arrLeft == 0 {
			return math.MaxInt64
		}

		if val, exist := cache[[2]int{i, arrLeft}]; exist {
			return val
		}

		curSum := 0
		res := math.MaxInt64
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			next := dfs(j+1, arrLeft-1)
			if next != math.MaxInt64 {
				res = min(res, max(curSum, next))
			}
		}

		cache[[2]int{i, arrLeft}] = res

		return res
	}

	return dfs(0, k)
}
func coinChange(coins []int, amount int) int {
	// brute force approach, recursive native
	var dfs func(am int) int 

	dfs = func(am int) int {
		if am == 0 {
			return 0
		}

		res := math.MaxInt32
		for _, coin := range coins {
			if am - coin >= 0 {
				res = min(res, 1 + dfs(am - coin))
			}
		}
		return res
	}

	minCoints := dfs(amount)
	if minCoints >= math.MaxInt32 {
		return -1
	}

	return minCoints
}

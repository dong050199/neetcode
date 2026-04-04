func coinChange(coins []int, amount int) int {
	var dfs func(am int) int 

	dfs = func(am int) int {
		if am == 0 {
			return 0
		}
		res := math.MaxInt32
		for _, coin := range coins {
			if coin <= am {
				res = min(res, 1 + dfs(am - coin))
			}
		}
		return res
	}

	minCoins := dfs(amount)
	if minCoins >= math.MaxInt32 {
		return -1
	}

	return minCoins
}

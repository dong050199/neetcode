func change(amount int, coins []int) int {
	sort.Ints(coins)
	var dfs func (i, a int) int 

	dfs = func (i, a int) int {
		if a == 0 {
			return 1
		}

		if i >= len(coins) {
			return 0
		}

		res := 0

		if a >= coins[i] {
			res = dfs(i + 1, a)
			res += dfs(i, a-coins[i])
		}			

		return res
	}

	// TODO: need to check if it impossible to make amunt later
	return dfs(0, amount)
}
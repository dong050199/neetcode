func change(amount int, coins []int) int {
	sort.Ints(coins)
	var dfs func (i, a int) int 

	memo := make(map[[2]int]int)

	dfs = func (i, a int) int {
		memoKey := [2]int{i, a}

		if _, exist := memo[memoKey]; exist {
			return memo[memoKey]
		}

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


		memo[memoKey] = res
		return res
	}

	// TODO: need to check if it impossible to make amunt later
	return dfs(0, amount)
}
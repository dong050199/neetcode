func shipWithinDays(weights []int, days int) int {
	var dfs func(i int, daysLeft int) int 
	dfs = func(i, daysLeft int) int {
		if i >= len(weights) {
			if daysLeft == 0 {
				return 0
			}
			return math.MaxInt64
		}

		if daysLeft == 0 {
			return math.MaxInt64
		}

		curWeight := 0
		res := math.MaxInt64
		for j := i; j < len(weights); j++ {
			curWeight += weights[j]
			nextW := dfs(j +1, daysLeft - 1)
			if nextW != math.MaxInt64 {
				res = min(res, max(curWeight, nextW))
			}
		}

		return res
	}

	return dfs(0, days)
}

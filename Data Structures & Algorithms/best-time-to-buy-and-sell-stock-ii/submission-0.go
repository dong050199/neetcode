func maxProfit(prices []int) int {
	var dfs func(i int, bought bool) int
	cache := make(map[cacheKey]int)
	dfs = func(i int, bought bool) int {
		if i >= len(prices) {
			return 0
		}

        if val, exist := cache[cacheKey{idx: i, bought: bought}]; exist {
            return val 
        }

		res := dfs(i+1, bought)

		if bought {
			if prices[i]+dfs(i+1, false) > res {
				res = prices[i] + dfs(i+1, false)
			}
		} else {
			if -prices[i]+dfs(i+1, true) > res {
				res = -prices[i] + dfs(i+1, true)
			}
		}

		cache[cacheKey{idx: i, bought: bought}] = res

		return res
	}

	return dfs(0, false)
}

type cacheKey struct {
	idx    int
	bought bool
}
func maxProfit(prices []int) int {
	res := 0

	l := 0
	for r := 0; r < len(prices); r++ {
		res = max(res, prices[r] - prices[l])
		if prices[r] < prices[l] {
			l = r
		}
	}

	return res
}

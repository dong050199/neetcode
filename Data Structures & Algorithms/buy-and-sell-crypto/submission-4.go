func maxProfit(prices []int) int {
	l := 0
	res := 0
	for r := 1 ; r < len(prices); r++ {
		res = max(res, prices[r] - prices[l])
		if prices[r] < prices[l] {
			l = r
			continue
		}
	}
	return res
}

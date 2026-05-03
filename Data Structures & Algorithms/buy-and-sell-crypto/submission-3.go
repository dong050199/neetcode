func maxProfit(prices []int) int {
	res := 0
	l := 0 
	for r := 1; r < len(prices); r++ {
		profit :=  prices[r] - prices[l]
		res = max(res, profit)
		if prices[r] < prices[l] {
			l = r
		}
	}
	return res
}

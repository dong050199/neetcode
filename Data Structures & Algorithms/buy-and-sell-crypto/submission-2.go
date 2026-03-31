func maxProfit(prices []int) int {
	// brute force approach is using nested for loop.
	profit := 0
	for i, price := range prices {
		for j := i + 1 ; j < len(prices) ; j++ {
			profit = max(profit , prices[j] - price)
		}
	}

	return profit
}

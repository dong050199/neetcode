func maxProfit(prices []int) int {
    var profit int
    for i, price := range prices {
        buyPrice := price
        sellPrice := 0
        // now need to find the efficent sell price
        for j := i+1; j < len(prices); j ++ {
            if prices[j] > sellPrice {
                sellPrice = prices[j]
            }
        }
        if (sellPrice - buyPrice) > profit {
            profit = sellPrice - buyPrice
        }
    }

    return profit
}

// let try with brute force first:
// [10,1,5,6,7,1]
// if we buy at day 0 -> find the max of [1,5,6,7,1] = 7 -> profit is 0 (because 10 > 7?)
// if we buy at day 1 -> find the max of [5,6,7,1] = 7 -> profit is 6 -> ok that is


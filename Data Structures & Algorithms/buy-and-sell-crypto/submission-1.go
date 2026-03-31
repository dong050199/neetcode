func maxProfit(prices []int) int { // this using two pointer
    maxP := 0
    minBuy := 100

    for _, sell := range prices {
        if sell - minBuy > maxP {
            maxP = sell - minBuy
        }
        if sell < minBuy {
            minBuy = sell
        }
    }
    return maxP
}

// max profit meaning we have to find the two things:
// the lowsest price so far: 
// the best profit so far:


// the prices:
// [10,1,5,6,7,1]
// if we buy at day 0 and sell at day 1 what happen?
// 10 - 1
// hmm, that we are lossing money. So need to 
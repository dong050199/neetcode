func maxProduct(nums []int) int {
    res := nums[0]
	
	curMax, curMin := 1, 1

	for _, num := range nums {
		prevMax, prevMin := curMax, curMin
		curMax = max(prevMax * num, prevMin * num, num)
		curMin = min(prevMax * num, prevMin * num, num)
		res = max(res, curMax)
	}

	return res
}

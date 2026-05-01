func maxProduct(nums []int) int {
    res := nums[0]
	prevMax, prevMin := 1, 1
	for _, num := range nums {
		curMax, curMin := prevMax, prevMin
		prevMax = max(num, num * curMax, num *curMin)
		prevMin = min(num, num * curMax, num *curMin)
		res = max(res, prevMax)
	}
	return res
}

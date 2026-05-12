func maxProduct(nums []int) int {
	res := 0
	curMax, curMin := 1, 1
	for _, num := range nums {
		preMax, preMin := curMax, curMin
		curMax = max(preMax * num, preMin * num, num)
		curMin = min(preMax * num, preMin * num, num) 
		res = max(res, curMax)
	}
	return res
}

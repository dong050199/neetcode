func maxProduct(nums []int) int {
	res := 1

	curMax, curMin := 1, 1
	for _, num := range nums {
		preMax, prevMin := curMax, curMin
		curMax = max(prevMax * num, prevMin * num, num)
		curMin = min(prevMax * num, prevMin * num, num) 
		res = max(res, curMax)
	}
	return res
}

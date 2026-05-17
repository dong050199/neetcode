func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}
	res := 0
	maxProd, minProd := 1, 1
	for i := 0; i < len(nums); i++ {
		curMax, curMin := maxProd, minProd
		maxProd = max(curMax * nums[i], curMin * nums[i], nums[i])
		minProd = min(curMax * nums[i], curMin * nums[i], nums[i])
		res = max(res, maxProd)
	}
	return res
}

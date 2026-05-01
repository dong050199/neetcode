func maxSubArray(nums []int) int {
    res := nums[0]
	prevMax := nums[0]

	for i := 1; i < len(nums); i++ {
		prevMax = max(nums[i], prevMax + nums[i])
		res = max(res, prevMax)
	}

	return res
}

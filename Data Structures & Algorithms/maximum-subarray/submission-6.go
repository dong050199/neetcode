func maxSubArray(nums []int) int {
	maxSub := nums[0]
	curSub := nums[0]

	for _, n := range nums[1:] {
		curSub = max(n, curSub + n)
		maxSub = max(maxSub, curSub)
	}

	return maxSub
}
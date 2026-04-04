func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

    maxSub := -10000
	curSub := 0

	for _, n := range nums {
		if n > curSub {
			curSub = n
		} else {
			curSub += n
		}

		maxSub = max(maxSub, curSub)
	}

	return maxSub
}
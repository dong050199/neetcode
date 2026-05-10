func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := 0
	maxSum := 0

	for _, num := range nums {
		curSum := maxSum
		maxSum = max(curSum + num, num)
		res = max(res, maxSum)
	}
    return res
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := -1001

	for _, num := range nums {
		curSum := maxSum
		maxSum = max(curSum + num, num)
	}
    return maxSum
}

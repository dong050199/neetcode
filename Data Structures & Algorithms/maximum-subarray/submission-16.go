func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1 ; i < len(nums); i++ {
		maxSum = max(maxSum + nums[i], nums[i])
	}
    return maxSum
}

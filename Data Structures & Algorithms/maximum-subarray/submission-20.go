func maxSubArray(nums []int) int {
    maxSum := 0

	for _, num := range nums {
		maxSum = max(maxSum, num)
		maxSum += num
	}

	return maxSum
}

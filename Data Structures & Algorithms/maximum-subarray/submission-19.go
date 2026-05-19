func maxSubArray(nums []int) int {
    maxSum := 0

	for _, num := range nums {
		maxSum += num
		maxSum = max(maxSum, num)
	}

	return maxSum
}

func maxSubArray(nums []int) int {
    maxSum := 0

	for _, num := range nums {
		maxSum = max(maxSum + num, num)
	}

	return maxSum
}

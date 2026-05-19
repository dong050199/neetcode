func maxSubArray(nums []int) int {
    maxSum := 0

	for _, num := range nums {
		maxSum += num
		if maxSum < num {
			maxSum = num
			continue
		}
	}

	return maxSum
}

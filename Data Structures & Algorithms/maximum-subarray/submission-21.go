func maxSubArray(nums []int) int {
    maxSum := 0

	for _, num := range nums {
		if maxSum < num {
			maxSum = num
			continue
		}
		maxSum += num
	}

	return maxSum
}

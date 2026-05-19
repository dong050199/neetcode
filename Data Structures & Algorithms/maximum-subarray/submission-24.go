func maxSubArray(nums []int) int {
    maxSum := 0
	res := 0

	for _, num := range nums {
		maxSum = max(maxSum + num, num)
		res = max(res, maxSum)
	}

	return res
}

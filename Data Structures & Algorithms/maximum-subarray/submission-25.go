func maxSubArray(nums []int) int {
    maxSum := -10000
	res := -100000

	for _, num := range nums {
		maxSum = max(maxSum + num, num)
		res = max(res, maxSum)
	}

	return res
}

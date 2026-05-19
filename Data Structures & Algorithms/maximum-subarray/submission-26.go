func maxSubArray(nums []int) int {
    maxSum := -1000
	res := -1000
	for _, num := range nums {
		maxSum = max(maxSum + num, num)
		res = max(res, maxSum)
	}
	return res
}

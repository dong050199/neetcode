func maxProduct(nums []int) int {
	res := nums[0]

	maxDp := 1
	minDp := 1

	for _, num := range nums {
		prevMax, prevMin := maxDp, minDp

		maxDp = max(num, prevMax * num, prevMin * num)
		minDp = max(num, prevMax * num, prevMin * num)
	}

	res = max(res, maxDp)

	return res
}

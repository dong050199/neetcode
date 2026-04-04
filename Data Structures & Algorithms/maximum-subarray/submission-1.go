func maxSubArray(nums []int) int {
	ma := nums[0]
	for i := 0; i < len(nums); i++ {
		mi := 0
		for j := i; j < len(nums); j++ {
			mi = mi + nums[j]
		}
		ma = max(ma, mi)
	}

	return ma
}

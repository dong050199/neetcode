func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ma := 0
	for i := 0; i < len(nums); i++ {
		mi := 0
		for j := i; j < len(nums); j++ {
			mi = mi + nums[j]
		}
		ma = max(ma, mi)
	}

	return ma
}

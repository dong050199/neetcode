func maxSubArray(nums []int) int {
	// brute force 
	maxV := -1001
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		for j := i; j < len(nums); j++ {
			if j > i {
				cur += nums[j]
			}
			maxV = max(maxV, cur)
		}
	}

	return maxV
}

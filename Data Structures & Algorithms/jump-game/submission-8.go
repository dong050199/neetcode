func canJump(nums []int) bool {
	l, r := 0, 0
	
	for r < len(nums) - 1 {
		mr := 0
		for i := l; i <= r; i ++ {
			mr = max(mr, nums[i] + i)
		}

		if mr <= r {
			return false
		}

		l = r + 1
		r = mr
	}

	return true
}
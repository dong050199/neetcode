func canJump(nums []int) bool {
    maxJump := 0

	for i := 0; i < len(nums); i++ {
		if maxJump < i {
			return false
		}
		maxJump = max(maxJump, i + nums[i])
	}

	return maxJump >= len(nums) - 1
}

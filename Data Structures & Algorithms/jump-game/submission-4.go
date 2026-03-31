func canJump(nums []int) bool {
	goal := len(nums) - 1 // last index

	for i := len(nums) - 2; i >= 0 ; i-- {
		if nums[i] + i >= goal {
			goal = i
		}
	}

	return goal == 0
}

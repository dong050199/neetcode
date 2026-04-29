func canJump(nums []int) bool {
    goal := len(nums)-1


	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= goal - i {
			goal = i
		}
	}

	return goal == 0
}

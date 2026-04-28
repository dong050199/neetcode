func canJump(nums []int) bool {
    goal := len(nums) - 1

	for g := len(nums) -2; g >= 0; g-- {
		if g + nums[g] >= goal {
			goal = g
		}
	}
	return goal == 0
}

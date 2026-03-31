func canJump(nums []int) bool {
    goal := len(nums) - 1

    for i := len(nums) - 2; i >= 0; i-- {
        if i+nums[i] >= goal {
            goal = i  // this index can reach the goal, so it becomes new goal
        }
    }

    return goal == 0
}
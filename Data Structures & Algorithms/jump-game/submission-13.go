func canJump(nums []int) bool {
    maxJump := 0
	for i, num := range nums {
		if maxJump < i {
			return false
		}
		maxJump = max(maxJump, num + i)
	}
	return true
}

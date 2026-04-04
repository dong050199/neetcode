func canJump(nums []int) bool {
	if nums[0] == 0 {
		return false
	}

	tmp := 1
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] == 0 {
			tmp +=1
			continue
		}

		if nums[i] >= tmp {
			tmp = 1
			continue
		} 
		return false
	}

	return true
}

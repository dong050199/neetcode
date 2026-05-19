func jump(nums []int) int {
    maxJump := 0
	res := 0
	for i, num := range nums {

		if i + num > maxJump {
			maxJump = i + num 
			if i == len(nums) - 1 {
				if maxJump >= 0 {
					return res
				}
			}
			res++
		}		
	}
	return res
}

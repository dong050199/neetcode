func productExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	cur := 1
	for i := 0; i < len(nums); i++ {
		cur = cur*nums[i]
		left[i] = cur
	}	

	cur = 1
	for i := len(nums)-1; i >= 0; i-- {
		cur = cur*nums[i]
		right[i] = cur
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res = append(res, right[i+1])
			continue
		}

		if i == len(nums) - 1 {
			res = append(res, left[i-1])
			continue
		}

		cur = left[i-1] * right[i+1]
		res = append(res, cur)
	}
	return res
}

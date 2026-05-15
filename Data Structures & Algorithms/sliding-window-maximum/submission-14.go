func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	q := []int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && q[0] < (i - k + 1) {
			q = q[1:]
		}

		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i - k + 1 >= 0 {
			res = append(res, nums[q[0]])
		}
	}

	return res
}

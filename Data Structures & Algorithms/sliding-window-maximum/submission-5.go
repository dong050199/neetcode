func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
	q := []int{}

	l, r := 0, 0
	for r < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		
		if q[0] < l {
			q = q[1:]
		}

		if r + 1 >= k {
			res = append(res, nums[q[0]])
			l++
		}
		r++
	}
	return res
}

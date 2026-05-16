func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}

	l := 0
	q := []int{}
	for r := 0; r < len(nums); r++ {
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		for len(q) > 0 && q[0] < r - k + 1 {
			q = q[1:]
			l++
		}

		q = append(q, r)

		if r >= k - 1 {
			res = append(res, nums[q[0]])
		}
	}



	return res
}

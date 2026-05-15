func maxSlidingWindow(nums []int, k int) []int {
	// monotonic dequeue
	dq := []int{}
	res := []int{}


	for i := 0; i < len(nums); i++ {
		for len(dq) > 0 && dq[0] < (i - k + 1) {
			dq = dq[1:]
		}

		for len(dq) > 0 && nums[i] > nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)

		if i - k + 1 >= 0 {
			res = append(res, dq[0])
		}
	}

	return res
}

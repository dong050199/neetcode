func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
	queue := []int{}
	l := 0
	for r := 0; r < len(nums); r++ {
		for len(queue) > 0 && nums[r] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, r)

		for len(queue) > 0 && queue[0] < l {
			queue = queue[1:]
		}

		if r >= k - 1 {
			res = append(res, nums[queue[0]])
			l++
		}
	}

	return res
}

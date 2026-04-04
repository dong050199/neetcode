func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	// brute force approach
	window := make([]int, k)
	for i := 0; i < len(nums) - k + 1; i ++ {
		copy(window,nums[i:i+k])
		max := maxOf(window)
		res = append(res, max)
	}

	return res
}

func maxOf(window []int) int {
	sort.Ints(window)
	return window[len(window)-1]
}



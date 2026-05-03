func maxArea(heights []int) int {
	var res int
	l, r := 0, len(heights) - 1
	maxLeft := 0
	maxRight := 0

	for l < r {
		maxLeft = max(maxLeft, heights[l])
		maxRight = max(maxRight, heights[r])
		res = max(res, min(maxLeft, maxRight) * (r - l))

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}


	return res
}

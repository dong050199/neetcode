func maxArea(heights []int) int {
	res := 0
	l, r := 0, len(heights) - 1

	maxLeft, maxRight := heights[l], heights[r]

	for l < r {
		res = max(res, (r - l) * min(maxLeft, maxRight))
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, heights[l])
		} else {
			r--
			maxRight = max(maxRight, heights[r])
		}
	}

	return res
}

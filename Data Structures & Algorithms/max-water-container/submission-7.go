func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	res := 0
	maxLeft, maxRight := heights[l], heights[r]
	for l < r {
		res = max(res , (r - l) * min(maxLeft, maxRight))
		
		if maxLeft > maxRight {
			r--
		} else {
			l++
		}

		maxRight = max(maxRight, heights[r])
		maxLeft = max(maxLeft, heights[l])
	}

	return res
}

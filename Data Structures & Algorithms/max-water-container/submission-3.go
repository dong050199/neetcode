func maxArea(heights []int) int {
	res := 0
	l, r := 0, len(heights) - 1
	maxLeft, maxRight := heights[l], heights[r]
	for l < r {
		res = max(res, (r - l) * min(maxLeft, maxRight))
		if heights[l] > heights[r] {
			r--
			maxRight = max(maxRight, heights[r])
		} else {
			l++
			maxLeft = max(maxLeft, heights[l])
		}
	}
	return res
}	

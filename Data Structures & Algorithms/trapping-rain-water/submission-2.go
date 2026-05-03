func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0

	l, r := 0, len(height) - 1
	maxLeft, maxRight := height[l], height[r]
	
	for l < r {
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, height[l])
			res+= maxLeft - height[l]
		} else {
			r--
			maxRight = max(maxRight, height[r])
			res+= maxRight - height[r]
		}
	}

	return res
}

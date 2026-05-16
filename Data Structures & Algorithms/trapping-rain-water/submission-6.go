func trap(height []int) int {
	l, r := 0, len(height) - 1
	res := 0
	maxLeft, maxRight := height[l], height[r]
	for l < r {
		if maxLeft < maxRight {
			l++
			if height[l] < maxLeft {
				res+= maxLeft - height[l]
			} else {
				maxLeft = height[l]
			}
		} else {
			r--
			if height[r] < maxRight {
				res+= maxRight - height[r]
			} else {
				maxRight = height[r]
			}
		}
	}

	return res
}

func trap(height []int) int {
	l, r := 0, len(height) - 1
	maxLeft, maxRight := height[l], height[r]
	res := 0
	for l < r {
		if height[r] > height[l] {
			if height[l] < maxLeft {
				res += maxLeft - height[l]
			} else {
				maxLeft = height[l]
			}
			l++
		} else {
			if height[r] < maxRight {
				res += maxRight - height[r]
			} else {
				maxRight = height[r]
			}
			r--
		}
	}
	return res
}

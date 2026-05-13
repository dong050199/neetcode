func trap(height []int) int {
	l, r := 0, len(height) - 1
	maxLeft, maxRight := height[l], height[r]

	res := 0

	for l < r {
		if maxLeft < maxRight {
			l++
			if maxLeft > height[l] {
				res += maxLeft - height[l]
				continue
			} 
			maxLeft = height[l]
		} else {
			r--
			if maxRight > height[r] {
				res += maxRight - height[r]
				continue
			}
			maxRight = height[r]
		}
	}

	return res
}

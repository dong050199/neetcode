func trap(height []int) int {
	// for each index we calculate the maximun water can be trapped and redude for bricks that filled.
	maxLeft := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i == 0 {
			maxLeft[i] = height[i]
			continue
		}
		if maxLeft[i-1] < height[i-1] {
			maxLeft[i] = height[i-1]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}

	maxRight := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			maxRight[i] = height[len(height)-1]
			continue
		}

		if height[i+1] > maxRight[i+1] {
			maxRight[i] = height[i+1]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	total := 0
	for i, h := range height {
		m := min(maxLeft[i], maxRight[i]) - h
		if m < 0 {
			m = 0
		}
		total += m
	}
	return total
}
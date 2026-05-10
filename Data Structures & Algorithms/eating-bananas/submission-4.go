func minEatingSpeed(piles []int, h int) int {
	l, r := 0, maxOf(piles)

	for l <= r {
		mid := l + (r - l) / 2

		time := 0 
		for _, pile := range piles {
			if pile <= mid {
				time++
			} else {
				if pile % mid == 0 {
					time += pile/mid
				} else {
					time += pile/mid + 1
				}
			}
		}

		if time > h {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func maxOf(piles []int) int {
	maxV := piles[0]
	for _, pile := range piles {
		maxV = max(maxV, pile)
	}
	return maxV
}

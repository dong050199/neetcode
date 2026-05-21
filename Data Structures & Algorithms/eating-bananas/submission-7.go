func minEatingSpeed(piles []int, h int) int {
	// search in piles space
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	l := 1
	r := maxPile
	for l <= r {
		mid := l + (r - l)/2

		time := 0
		for _, pile := range piles {
			if pile < mid {
				time++
			} else {
				if pile % mid == 0 {
					time += pile / mid
				} else {
					time += pile / mid + 1
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

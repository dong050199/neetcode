func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, maxOf(piles)

	for lo < hi {
		mid := lo + (hi - lo)/2
		
		times := 0
		for _, pile := range piles {
			times += (pile + mid - 1) / mid
		}

		if times <= h {
            hi = mid
        } else {
            lo = mid + 1
        }
	}

	return lo
}

func maxOf(piles []int) int {
	res := piles[0]

	for _, pile := range piles {
		if pile > res {
			res = pile
		}
	}

	return res
}

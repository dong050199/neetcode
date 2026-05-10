func minEatingSpeed(piles []int, h int) int {
	l, r := 1, maxOf(piles)

	for l <= r {
		mid := l + (r-l)/2
		time := 0
		for _, p := range piles {
			if p <= mid {
				time++
			} else {
				if p % mid == 0 {
					time += p / mid
				} else {
					time += p / mid + 1
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

func maxOf(req []int) int {
	res := req[0]
	for _, r := range req {
		res = max(res, r)
	}
	return res
}
func carPooling(trips [][]int, capacity int) bool {
	// this quite same with interval
	// we need to sort trips by from first
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

    for i := 0; i < len(trips); i++ {
        curPass := trips[i][0]
        for j := 0; j < i; j++ {
            if trips[j][2] > trips[i][1] {
                curPass += trips[j][0]
            }
        }

        if curPass > capacity {
            return false
        }
    }
    return true
}
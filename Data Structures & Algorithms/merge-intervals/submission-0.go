func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort intervals based on the start 
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// check overlap and merge
	res := [][]int{intervals[0]}
	for i := 1 ; i < len(intervals); i++ {
		cur := res[len(res) - 1]
		// if overlap then merge and update cur 
		if intervals[i][0] <= cur[1] {
			cur[0] = min(cur[0], intervals[i][0])
			cur[1] = max(cur[1], intervals[i][1])
			res[len(res) - 1] = cur
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

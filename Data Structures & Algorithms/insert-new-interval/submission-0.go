func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	for i, iv := range intervals {
		if newInterval[1] < iv[0] {
			res = append(res, newInterval)
			return append(res, intervals[i:]...) 
		} else if iv[1] < newInterval[0] {
			res = append(res, iv)
		} else {
			newInterval[0] = min(newInterval[0], iv[0])
			newInterval[1] = max(newInterval[1], iv[1])
		}
	}

	return append(res, newInterval)
}

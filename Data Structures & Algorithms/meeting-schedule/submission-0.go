/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) <= 1 {
		return true
	}
	// first we need to sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		prev := intervals[i-1]

		if cur.start < prev.end {
			return false
		}
	}

	return true
}

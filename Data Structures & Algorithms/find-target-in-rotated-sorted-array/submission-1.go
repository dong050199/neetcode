func search(nums []int, target int) int {
	// need to find the pivot
	l, r := 0, len(nums) - 1
	for l < r {
		m := l + (r - l) / 2
		if nums[m] < nums[r] {
			r = m 
			continue
		}
		l = m + 1
	}
	pivot := l

	// Determine which half to search (use <= for boundaries)
    l, r = 0, len(nums)-1
    if target >= nums[pivot] && target <= nums[r] {
        l = pivot
    } else {
        r = pivot - 1
    }
    
    // Standard binary search with l <= r
    for l <= r {
        m := l + (r-l)/2
        if nums[m] == target {
            return m
        }
        if nums[m] < target {
            l = m + 1  // Always m + 1
        } else {
            r = m - 1  // Always m - 1
        }
    }
    
    return -1
}

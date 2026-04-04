func search(nums []int, target int) int {
	// need to find the pivot
	l, r := 0, len(nums) - 1
	for l < r {
		m := l + (r - l) / 2
		if nums[m] < nums[r] {
			r = m - 1
			continue
		}
		l = m
	}
	// if the target in the nums[l] and nums[l - 1]
	min := nums[l]
	max := nums[l]
	if l == 0 {
		max = nums[len(nums) - 1]
	} else {
		max = nums[l - 1]
	}

	// min and max
	if target > max && target < min {
		return - 1
	}

	if nums[l] < target && target < nums[len(nums) - 1] {
		r = len(nums) - 1
	} else { 
		l , r = 0, l - 1
	}

	for l < r {
		m := l + (r - l) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1 
			continue
		}
		l = m
	}

	return -1
}

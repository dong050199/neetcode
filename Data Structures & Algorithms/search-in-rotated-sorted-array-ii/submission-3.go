func search(nums []int, target int) bool {
	// this problem must be solve by using binary search, for sure
	// first we need to find peak of the array where the order of array was changed
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r - l) / 2
		// [2,3,4,5,6,1] 
		// [4,5,6,1,2,3] 
		if nums[mid] > nums[r] {
			l = mid +1
		} else {
			r = mid - 1
		}
	}

	peak := l
	// now we have two sub arr
	if nums[0] > target {
		l = peak + 1
		r = len(nums) - 1
	} else {
		l = 0
		r = peak
	}

	for l <= r {	
		mid := l + (r - l)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid +1
		}
	}

	return false
}

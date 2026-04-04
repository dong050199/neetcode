func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums) - 1


	for lo < hi {
		mid := lo + (hi - lo)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid
		}
	}

	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	for len(nums) > 0 {
		mid := len(nums) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			nums = nums[:mid]
		} else {
			nums = nums[mid+1:]
		}
	}

	return -1
}

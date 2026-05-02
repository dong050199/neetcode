func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)
	res := 1
	cal := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]  {
			continue
		}

		if nums[i] == nums[i-1] + 1 {
			cal++
		} else {
			cal = 1
		}

		res = max(res, cal)
	}

	return res
}

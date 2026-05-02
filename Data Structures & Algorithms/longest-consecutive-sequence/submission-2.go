func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	res := 0
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

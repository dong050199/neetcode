func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j] 
	})

	res := 1
	cal := 1

	for i := 1; i < len(nums); i ++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == nums[i-1] + 1 {
			cal++
			res = max(res, cal)
		} else {
			cal = 1
		}
	}

	return res
}

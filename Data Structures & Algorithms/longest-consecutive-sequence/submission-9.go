func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	res := 1
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] == nums[i-1] + 1 {
			cur++
			res = max(res, cur)
			continue
		}
		cur = 1
	}

	return res
}

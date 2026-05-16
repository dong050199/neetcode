func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	res := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if nums[i] != nums[i-1] + 1 {
			count = 1
			continue
		}
		count++
		res = max(res, count)
	}

	return res
}

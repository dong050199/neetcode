func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	mp := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		l, r := i + 1, len(nums) - 1
		for l < r {
			if nums[i] + nums[l] + nums[r] == 0 {
				mp[[3]int{nums[i], nums[l], nums[r]}] = true
				break
			}
			if nums[i] + nums[l] + nums[r] >= 0 {
				r--
				continue
			}
			if nums[i] + nums[l] + nums[r] < 0 {
				l++
				continue
			}
		}
	}

	for k := range mp {
		res = append(res, k[:])
	}

	return res
}

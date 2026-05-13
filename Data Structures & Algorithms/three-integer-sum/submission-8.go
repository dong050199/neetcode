func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resMap := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		l, r := i + 1, len(nums) - 1
		for l < r {
			if nums[l] + nums[r] + nums[i] == 0 {
				resMap[[3]int{nums[i], nums[l], nums[r]}] = true
			}

			if nums[l] + nums[r] + nums[i] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	res := [][]int{}
	for k := range resMap {
		res = append(res, k[:])
	}
	return res
}

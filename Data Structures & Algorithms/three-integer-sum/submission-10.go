func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resMp := make(map[[3]int]bool)
	for i := 0; i < len(nums)-1; i++ {
		l, r := i + 1, len(nums)-1
		for l < r {
			if nums[l] + nums[r] + nums[i] == 0 {
				resMp[[3]int{nums[i], nums[l], nums[r]}] = true
				break
			}

			if nums[l] + nums[r] > nums[i] {
				r--
			} else {
				l++
			}
		}
	}

	res := [][]int{}
	for k := range resMp {
		res = append(res, k[:])
	}

	return res
}

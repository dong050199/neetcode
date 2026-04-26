func subsets(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < (1 << len(nums)); i ++ {
		subset := []int{}
		for j := 0; j < len(nums); j++ {
			if (i & (1 << j)) != 0 {
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}

	return res
}
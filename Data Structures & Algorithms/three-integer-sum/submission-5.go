func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		left, right := i + 1, len(nums) - 1
		for left < right {
			if nums[left] + nums[right] + nums[i] == 0 {
				res = append(res, []int{nums[i],nums[left],nums[right]})
				break
			}

			if nums[left] + nums[right] + nums[i] >= 0 { 
				right--
				continue
			}


			if nums[left] + nums[right] + nums[i] < 0 {
				left++
				continue
			}
		}
	}

	mp := make(map[[3]int]bool)
	newRes := [][]int{}
	for _, r := range res {
		if _, exist := mp[[3]int{r[0],r[1],r[2]}]; exist {
			continue
		}
		mp[[3]int{r[0],r[1],r[2]}] = true
		newRes = append(newRes, r)
	}

	return newRes
}
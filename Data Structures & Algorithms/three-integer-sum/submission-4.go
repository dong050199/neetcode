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

	return res
}
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j] 
    })

	var res [][]int
	mp := make(map[[3]int]bool)
	for i, num := range nums {
		left, right := i + 1, len(nums) - 1 

		for left < right {
			if nums[left] + nums[right] == - num {
				mp[[3]int{num, nums[left], nums[right]}] = true
			}

			if nums[left] + nums[right] > - num {
				right--
			} else {
				left++
			}
		}
	}

	for k := range mp {
		res = append(res, k[:])
	}

	return res
}

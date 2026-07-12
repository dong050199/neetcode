func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				switch {
				case sum == target:

					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left+1] == nums[left] {
						left++
					}

					for left < right && nums[right-1] == nums[left] {
						right--
					}

					left++
					right--
				case sum > target:
					right--

				case sum < target:
					left++

				}
			}
		}
	}
	return res
}
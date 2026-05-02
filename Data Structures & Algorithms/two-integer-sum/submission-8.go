func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[target - nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, exist := mp[nums[i]]; exist {
			if i == mp[nums[i]] {
				continue
			}
			return []int{i, mp[nums[i]]}
		}
	}

	return []int{}
}

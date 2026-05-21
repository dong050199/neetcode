func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i] 
		if val, exist := mp[diff]; exist {
			if val == i {
				continue
			}
			return []int{i, val}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i, num := range nums {
		diff := target - num
		if val, exist := mp[diff]; exist {
			return []int{i, val}
		}
	}

	return []int{}
}

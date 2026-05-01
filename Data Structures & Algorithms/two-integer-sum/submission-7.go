func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range nums {
		mp[target - num] = i
	}

	for i, num := range nums {
		if idx, exist := mp[num]; exist {
			if i == idx {
				continue
			}
			return []int{i, idx}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range nums {
		mp[num] = i
	}

	for i, num := range nums {
		diff := target - num

		if j, exist := mp[diff]; exist {
			return []int{i, j}
		}
	}

	return []int{}
}

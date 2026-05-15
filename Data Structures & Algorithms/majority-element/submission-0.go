func majorityElement(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}

	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
		if mp[num] > len(nums)/2 {
			return num
		}
	}
	
	return 0
}

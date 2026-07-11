func firstMissingPositive(nums []int) int {
	// cannot to sort because sort using O(nlogn) time complexity?
	mp := make(map[int]int)
	less := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}
		mp[nums[i]] = i
		less = min(less, nums[i])
	}

    if less > 1 {
        return 1
    }

	for i := 0; i < len(nums); i++ {
		if _, exist := mp[less+1]; exist {
			less++
		} else {
			return less + 1
		}
	}

	return 0
}
func findDuplicate(nums []int) int {
    mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
		if mp[n] > 1 {
			return n
		}
	}
	return -1
}

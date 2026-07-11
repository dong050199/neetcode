func subarraySum(nums []int, k int) int {
	prefix := 0
	prefixMap := make(map[int]int)
    prefixMap[0] = 1
	count := 0
	for _, num := range nums {
		prefix += num
		if c, exist := prefixMap[prefix-k]; exist {
			count += c
		}
		prefixMap[prefix]++
	}
	return count
}
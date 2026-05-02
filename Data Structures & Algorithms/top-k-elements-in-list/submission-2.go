func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	sl := [][]int{}

	for k, v := range mp {
		sl = append(sl, []int{k, v})
	}

	sort.Slice(sl, func(i, j int) bool {
    	return sl[i][1] > sl[j][1]
	})

	res := []int{}
	for _, s := range sl[:k] {
		res = append(res, s[0])
	}

	return res
}

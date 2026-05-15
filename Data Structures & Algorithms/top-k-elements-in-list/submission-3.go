func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	feq := [][]int{}
	for k, v := range mp {
		feq = append(feq, []int{k, v})
	}

	sort.Slice(feq, func(i, j int) bool {
    	return feq[i][1] > feq[j][1]
	})

	res := []int{}
	for i, f := range feq {
		if i >= k {
			break
		}
		res = append(res, f[0])
	}

	return res
}

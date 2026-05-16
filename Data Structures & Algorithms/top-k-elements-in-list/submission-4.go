func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	freqSlice := [][]int{}

	for k, v := range freqMap {
		freqSlice = append(freqSlice, []int{k, v})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
    	return freqSlice[i][1] > freqSlice[j][1]
	})

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, freqSlice[i][0])
	}

	return res
}

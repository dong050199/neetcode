func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	arr := [][2]int{}
	for k, v := range mp {
		arr = append(arr, [2]int{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	var res []int

	for i := 0; i < k; i ++ {
		res = append(res, arr[i][0])
	}

	return res
}

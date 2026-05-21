func topKFrequent(nums []int, k int) []int {
	// sort //heap?
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	arr := [][]int{}
	for k, v := range mp {
		arr = append(arr, []int{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
    	return arr[i][1] > arr[j][1]
	})

	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res, arr[i][0])
	}

	return res
}

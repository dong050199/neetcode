
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	fmt.Println(count)

	buckets := make([][]int, len(nums)+1)
	for k, v := range count {
		buckets[v] = append(buckets[v], k)
	}

	fmt.Println(buckets)

	output := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, v := range buckets[i] {
			output = append(output, v)
			if len(output) == k {
				return output
			}
		}
	}

	return output
}
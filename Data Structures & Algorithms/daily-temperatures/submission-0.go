func dailyTemperatures(temperatures []int) []int {
	res := []int{}

	for i := 0; i < len(temperatures); i ++ {
		count := 0
		for j := i; j<len(temperatures); j ++ {
			if temperatures[j] <= temperatures[i] {
				continue
			}
			count = j - i 
			break
		}
		res = append(res, count)
	}

	return res
}


// Input: temperatures = [30,38,30,36,35,40,28]
// Output: [1,4,1,2,1,0,0]
// brute force approach can be:
// iteral for all day temperator and then check the rest and count for warmer days.
// this can be O(n^2)

// 
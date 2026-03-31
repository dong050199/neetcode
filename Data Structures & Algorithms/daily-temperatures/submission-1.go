func dailyTemperatures(temperatures []int) []int {
	res := []int{}
	// brute foce approach

	for i, temp := range temperatures {
		cal := 0
		for j := i + 1; j < len(temperatures); j ++ {
			if temp < temperatures[j] {
				cal = j - i
				break
			}
		}

		res = append(res, cal)
	}

	return res
}

//[30,38,30,36,35,40,28]
// 

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	stack := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = max(res[i], stack[len(stack)-1] - i)
		}
		
		stack = append(stack, i)
	}

	return res
}

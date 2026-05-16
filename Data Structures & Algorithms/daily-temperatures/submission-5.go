func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temerature[i] > temperatures[stack[len(stack)-1]] {
			stack = stack[len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	return res
}

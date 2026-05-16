func dailyTemperatures(temperatures []int) []int {
	// brute for approach will be iterate throught all days from beginging and calculate the new arr,
	// this approach will be O(n^2)
	// not optimise
	stack := []int{} // store index of the warmest day that so far we have seen.
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}	

		stack = append(stack, i)
	}

	return res
}

func calPoints(operations []string) int {
	stack := []int{}

	for i := 0; i < len(operations); i++ {
		cur := operations[i]
		switch {
		case cur == "+":
			prev1 := stack[len(stack)-1]
			prev2 := stack[len(stack)-2]
			stack = append(stack, prev1+prev2)
		case cur == "D":
			prev := stack[len(stack)-1]
			stack = append(stack, 2*prev)
		case cur == "C":
			stack = stack[:len(stack)-1]
		default:
			curNum, _ := strconv.Atoi(cur)
			stack = append(stack, curNum)
		}
	}

	res := 0
	for _, score := range stack {
		res += score
	}

	return res
}
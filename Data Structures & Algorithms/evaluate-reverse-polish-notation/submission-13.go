func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "*" || tokens[i] == "-" || tokens[i] == "/" || tokens[i] == "+" {
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			first :=stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tokens[i] == "*" {
				stack = append(stack, first * second)
			}

			if tokens[i] == "/" {
				stack = append(stack, first / second)
			}

			if tokens[i] == "+" {
				stack = append(stack, first + second)
			}

			if tokens[i] == "-" {
				stack = append(stack, first - second)
			}
			continue
		} 

		num, _ := strconv.Atoi(tokens[i])
		stack = append(stack, num)
	}

	return stack[0]
}

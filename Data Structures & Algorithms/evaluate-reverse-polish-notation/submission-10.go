func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		// Check if token is an operator
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// Pop two operands (note the order!)
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var result int
			switch token {
			case "+":
				result = first + second
			case "-":
				result = first - second
			case "*":
				result = first * second
			case "/":
				result = first / second
			}
			stack = append(stack, result)
		} else {
			// It's a number
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
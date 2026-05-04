func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "*" || token == "-" || token == "/" || token == "+" {
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			first := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var res int
			if token == "+" {
				res = first + second
			}

			if token == "-" {
				res = first - second
			}

			if token == "/" {
				res = first / second
			}

			if token == "*" {
				res = first * second
			}
			stack = append(stack, res)

		} else {
			tmp , _ := strconv.Atoi(token)
			stack = append(stack, tmp)
		}
	}
	return stack[0]
}

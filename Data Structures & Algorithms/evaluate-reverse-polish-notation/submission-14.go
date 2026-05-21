func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			number, _ := strconv.Atoi(tokens[i])
			stack = append(stack, number)
			continue
		}

		first := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		second := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if tokens[i] == "+" {
			stack = append(stack, first + second)
		}
		
		if tokens[i] == "-" {
			stack = append(stack, second - first)
		}

		if tokens[i] == "*" {
			stack = append(stack, second * first)	
		}
		
		if tokens[i] == "/" {
			stack = append(stack, second / first)
		}
	}

	return stack[0]
}

func evalRPN(tokens []string) int {
	stack := []int{}
	
	for _, t := range tokens {
		if isOperator(t) {
			// if operator just pop 2 last number from stack and do the operation
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // pop two last numbers from stack
			switch {
				case t == "*":
					stack = append(stack, a * b) 
				case t == "/":
					stack = append(stack, b / a) 
				case t == "+":
					stack = append(stack, b + a) 
				case t == "-":
					stack = append(stack, b - a) 
			}
		} else {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num) 
		}
	}

	return stack[0]
}

func isOperator(s string) bool {
	if s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}
	return false
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		if !isOperator(t) {
			tmp, _ := strconv.Atoi(t)
			stack = append(stack, tmp)
			continue
		}

		if t == "+" {
			cal := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, cal)
			continue
		}


		if t == "-" {
			cal :=  stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, cal)
			continue
		}

		if t == "*" {
			cal := stack[len(stack)-1]  * stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, cal)
			continue
		}

		if  t == "/" {
			cal := stack[len(stack)-1] / stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, cal)
			continue
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

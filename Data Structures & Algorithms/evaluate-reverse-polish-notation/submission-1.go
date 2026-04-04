func evalRPN(tokens []string) int {
	stack := []int{}

	res, _ := strconv.Atoi(tokens[0]) 

	for i := 0; i < len(tokens); i++ {
		if isOperator(tokens[i]) {
			switch {
				case tokens[i] == "*":
					res = res * stack[len(stack)-1] 
				case tokens[i] == "/":
					res = res / stack[len(stack)-1] 
				case tokens[i] == "+":
					res = res + stack[len(stack)-1] 
				case tokens[i] == "-":
					res = res - stack[len(stack)-1] 
			}
		} else {
			num, _ := strconv.Atoi(tokens[i]) 
			stack = append(stack, num)
		}
	}
	return res
}

func isOperator(s string) bool {
	if s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}
	return false
}

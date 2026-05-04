func evalRPN(tokens []string) int {
	first, _ := strconv.Atoi(tokens[0])
	tokens = tokens[1:]
	for len(tokens) > 1 {
		second, _ := strconv.Atoi(tokens[0])
		if tokens[1] == "+" {
			first = first + second
		}

		if tokens[1] == "-" {
			first = first - second
		}

		if tokens[1] == "*" {
			first = first * second
		}

		if tokens[1] == "/" {
			first = first / second
		}

		tokens = tokens[1:]
	}

	return first
}

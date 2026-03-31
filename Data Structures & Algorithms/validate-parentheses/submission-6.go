func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	var stack []rune

	for _, c := range s {
		// if open parentheses then put it to stack
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack , c)
			continue
		} 

		if len(stack) == 0 {
			return false
		}

		cur := stack[len(stack)-1]

		if cur != convert(c) {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func convert(c rune) rune {
	if c == ')' {
		return '('
	}

	if c == ']' {
		return '['
	}

	if c == '}' {
		return '{'
	}

	return '}'
} 

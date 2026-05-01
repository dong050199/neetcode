func isValid(s string) bool {
	stack := []byte{}


	if len(s) % 2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		} 

		if len(stack) == 0 {
			return false
		}

		if s[i] == ')' {
			cur := stack[len(stack)-1] 
			if cur == '(' {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		} 

		if s[i] == ']' {
			cur := stack[len(stack)-1] 
			if cur == '[' {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		} 

		if s[i] == '}' {
			cur := stack[len(stack)-1] 
			if cur == '{' {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		} 
	}
	return len(stack) == 0
}

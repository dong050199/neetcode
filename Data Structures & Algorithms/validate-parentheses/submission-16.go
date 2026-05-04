func isValid(s string) bool {
    stack := []byte{}

	mp := make(map[byte]byte)
	mp['}'] = '{'
	mp[']'] = '['
	mp[')'] = '('

	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] == mp[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}	

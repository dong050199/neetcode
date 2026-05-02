func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

	stack := []byte{}

	mp := make(map[byte]byte)
	mp['}'] = '{'
	mp[')'] = '('
	mp[']'] = '['

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if _, exist := mp[s[i]]; exist {
			cur := stack[len(stack)-1]
			if cur == mp[s[i]] {
				stack = stack[:len(stack)-1]
				continue
			}
		}
	}

	return true
}

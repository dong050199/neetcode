func decodeString(s string) string {
	stack := []byte{}
	numStack := []int{}
	num := 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		switch {
			case cur == ']':	
				// start calculate here
				prevChar := []byte{}
				for len(stack) > 0 && stack[len(stack)-1] != '[' {
					prevChar = append(prevChar, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 && stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				}

				tmp := []byte{}
				for i := len(prevChar)-1; i >= 0; i-- {
					tmp = append(tmp, prevChar[i]) 
				}
				prevChar = tmp

				// revert prevChar
				prevNum := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]

				str := []byte{}
				for i := 0; i < prevNum; i++ {
					str =append(str, prevChar...)
				}
				stack = append(stack, []byte(str)...)
			case cur == '[':
				numStack = append(numStack, num)
				stack = append(stack, cur)
				num = 0
			case unicode.IsDigit(rune(cur)):
				num = num * 10 + int(cur - '0')
			default:
				stack = append(stack, cur)
		}	
	}
	return string(stack)
}

func isValid(s string) bool {
    mp := map[rune]rune{'}':'{',']':'[',')':'('}
    stack := []rune{}

    if len(s) <= 1 {
        return false
    }

    for _, c := range s {
        if _, exist := mp[c]; !exist {
            stack = append(stack, c) // if it open bracket then put to stack
            continue
        }
        if stack[len(stack)-1] != mp[c] {
            return false
        }
        // need to pop the open bracket out of stack
        stack = stack[0:len(stack)-1] // pop open bracket our of stack
    }

    return len(stack) == 0
}
// for this problem we can use stack.
// the valid string must start with open brackets and if the close brackets appear means 
// the last element on stack mist be the valid one with it
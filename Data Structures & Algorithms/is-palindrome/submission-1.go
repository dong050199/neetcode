func isPalindrome(s string) bool {
   var arr []string

    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            arr = append(arr, string(unicode.ToLower(r)))
        }
    }
    
    if len(arr) == 0 {
        return true
    }

    for i := 0 ; i < len(s)/2; i ++ {
        if arr[i] != arr[len(arr) - 1 - i] {
            return false
        }
    }

    return true
}

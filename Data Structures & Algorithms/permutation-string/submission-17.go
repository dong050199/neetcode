func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    s1Arr := [26]int{}
    s2Arr := [26]int{}
    for i := 0; i < len(s1); i++ {
        s1Arr[s1[i] - 'a']++
        s2Arr[s2[i] - 'a']++
    }

    l := 0
    for r := len(s1); r < len(s2); r++ {
        if s1Arr == s2Arr {
            return true
        }
        s2Arr[s2[r] - 'a']++
        s2Arr[s2[l] - 'a']--
        l++
    }

    return s1Arr == s2Arr
}

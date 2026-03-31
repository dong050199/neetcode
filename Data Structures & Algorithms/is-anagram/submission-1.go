func isAnagram(s string, t string) bool {
    sM := make(map[string]int)
    tM := make(map[string]int)

    if len(s) != len(t) {
        return false
    }

    for _, c := range strings.Split(s, "") {
       sM[c]++
    }

    for _, c := range strings.Split(t, "") {
       tM[c]++
    }

    for _, c := range strings.Split(t, "") {
        if sM[c] != tM[c] {
            return false
        }
    }

    return true
}

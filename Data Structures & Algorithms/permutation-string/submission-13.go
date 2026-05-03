func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    for l, r := 0, len(s1) - 1; r < len(s2); l, r = l + 1, r + 1 {
        if isPermutation(s1, s2[l:r+1]) {
            return true
        }
    }

    return false
}

func isPermutation(s1, s2 string) bool {
    mp := make(map[byte]int)
    
    for i := 0; i < len(s1); i++ {
        mp[s1[i]]++
    }

    for i := 0; i < len(s2); i++ {
        mp[s2[i]]--
        if mp[s2[i]] == 0 {
            delete(mp, s2[i])
        }
    } 

    return len(mp) == 0
}

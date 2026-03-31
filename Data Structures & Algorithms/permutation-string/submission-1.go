func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    mp := make(map[byte]int)
    for i := 0; i < len(s1); i ++ {
        mp[s1[i]]++
    }

    for i:= 0; i <= len(s2) - len(s1); i++ {
        mpW := make(map[byte]int)    
        for j := i; j < i + len(s1); j++ {
            mpW[s2[j]]++
        }

        if !mapEqual(mp, mpW) {
            continue
        }
        return true
    }

    return false
}

func mapEqual(m1, m2 map[byte]int) bool {
    if len(m1) != len(m2) {
        return false
    }
    for k, v := range m1 {
        if m2[k] != v {
            return false
        }
    }
    return true
}
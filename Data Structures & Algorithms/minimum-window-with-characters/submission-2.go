func minWindow(s string, t string) string {
    minf := 0
    res := []int{}

    mp := make(map[byte]int)
    for i:=0; i<len(t); i ++ {
        mp[t[i]]++
    }

    for i := 0; i < len(s); i ++ {
        mpW := make(map[byte]int)
        for j := i; j<len(s); j ++ {
            mpW[s[j]]++
            if !isValidSubString(mp, mpW) {
                continue
            }

            if minf == 0 {
                minf = j - i + 1
            }
            
            if minf >= j - i + 1 {
                minf = j - i + 1
                res = []int{i, j}
            }
        }
    }

    if len(res) == 0 {
        return ""
    }

    return s[res[0]:res[1]+1]
}

func isValidSubString(m1, m2 map[byte]int) bool {
    for k, v := range m1 {
        if m2[k] < v {
            return false
        }
    }

    return true
}

// brute force approach can be
// check all possible substring from s, if substring contain all characters from t then update min
// optimise it?

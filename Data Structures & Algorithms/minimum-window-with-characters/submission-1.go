func minWindow(s string, t string) string {
    min := 0
    mp := make(map[byte]int)
    for i:=0; i<len(t); i ++ {
        mp[t[i]]++
    }

    mpW := make(map[byte]int)
    for i:=0; i<len(s); i++ {
    }

    return ""
}


// brute force approach can be
// check all possible substring from s, if substring contain all characters from t then update min
// optimise it?

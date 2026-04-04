func lengthOfLongestSubstring(s string) int {
    max :=0
    for i := 0; i < len(s); i ++ { // s[i] is uint8 0-255
        m := make(map[uint8]bool)
        counter := 0
        for j := i; j < len(s); j ++ {
            if !m[s[j]] {
                counter++
                m[s[j]] = true
                continue
            }
            // update max
            if counter > max {
                max = counter
            }
            break
        }
    }

    return max

}
func characterReplacement(s string, k int) int {
    l := 0
    res := 0
    mp := make(map[byte]int)
    maxFreq := 0

    for r := 0; r < len(s); r++ {
        mp[s[r]]++
    
        if mp[s[r]] > maxFreq {
            maxFreq = mp[s[r]]
        }
        for (r - l + 1) - maxFreq > k {
            mp[s[l]]--
            l++
        }
		res = max(res,  r - l + 1)
    }
    return res
}
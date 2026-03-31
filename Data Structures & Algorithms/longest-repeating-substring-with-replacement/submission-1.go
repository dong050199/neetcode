func characterReplacement(s string, k int) int {
    res := 0
    for i := 0; i< len(s); i++ {
        mp := make(map[byte]int)
        maxf := 0
        for j := i ; j < len(s); j ++ {
            mp[s[j]]++
            maxf = max(mp[s[j]], maxf)
            if (j - i + 1) - maxf <= k {
                res = max(res, (j - i + 1))
            }
        }
    }
    return res
}

// sliding window technique
// Solution 1: brute force
// try every possible substring at every index
// AAABABB , k = 1
// index = 0: from [A] [AA] [AAA] [AAAB] -> 1; [AAABA] -> 1; [AAABAB] -> 2(invalid) -> max = 5 
// index = 1: from [A] [AA] [AAB] [AABA] [AABAB] -> 2 (invalid) -> max = 4.
// ....
// invalid when: (length of sub string) - (max occurent character) > k
// max lenth = sub index
// try it ...
func lengthOfLongestSubstring(s string) int {
    l, length := 0, 0
    mp := make(map[byte]int) // map with key is character (to check exist) and value is index of character to check length.
    for r := 0; r < len(s); r ++ {
        if idx, exist := mp[s[r]]; exist {
            // if character already on map, then update left index to the current index of character + 1 (next character).
            l = idx + 1
        }
        // update index.
        mp[s[r]] = r
        length = max(length, r - l + 1)
    }
    return length
}
// technically, we can use the window sliding in this problem
// for example we have a window [0,0] means we have two characters in this window
// [zxyzxyz] -> [z] length = 1
// [zxyzxyz] -> [zx] length = 2
// [zxyzxyz] -> [zxy] length = 3
// [zxyzxyz] -> z[xyz] length = 3
// [zxyzxyz] -> zx[yzx] length = 3
// [zxyzxyz] -> zxy[zxy] length = 3
// ...
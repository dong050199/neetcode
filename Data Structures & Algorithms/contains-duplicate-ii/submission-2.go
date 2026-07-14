func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int][]int)
    for i, num := range nums {
        mp[num] = append(mp[num], i)
        if len(mp[num]) > 1 {
            last := mp[num][len(mp[num])-2]
            if i - last <= k {
                return true
            }
        }
    }

    return false
}

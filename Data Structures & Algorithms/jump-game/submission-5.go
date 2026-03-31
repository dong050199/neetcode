func canJump(nums []int) bool {
    l, r := 0, 0
    for r < len(nums)-1 {
        farthest := 0
        for i := l; i <= r; i++ {
            farthest = max(farthest, i+nums[i])
        }
        if farthest <= r {
            return false  // window didn't expand → stuck!
        }
        l = r + 1
        r = farthest
    }
    return true
}
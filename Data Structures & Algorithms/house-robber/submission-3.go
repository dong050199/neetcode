func rob(nums []int) int {
    ma := 0

    if len(nums) < 2 {
        return 0
    }

    if len(nums) == 2 {
        return nums[1] + nums[2]
    }

    for i, n := range nums {
        for j, m := range nums {
            if i == j + 1 || i == j - 1 || i == j {
                continue
            }
            total := m + n 
            ma = max(ma, total)
        }
    }

    return ma
}
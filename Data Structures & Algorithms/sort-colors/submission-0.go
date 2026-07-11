func sortColors(nums []int)  {
    r := 0
    for i := range nums {
        if nums[i] == 0 {
            nums[i], nums[r] = nums[r], nums[i]
            r++
        }
    }

    w := r
    for i := range nums {
        if nums[i] == 1 {
            nums[i], nums[w] = nums[w], nums[i]
            w++
        }
    }

    b := w
    for i := range nums {
        if nums[i] == 2 {
            nums[i], nums[b] = nums[b], nums[i]
            b++
        }
    }
}
func removeDuplicates(nums []int) int {
    // using two pointer
    l := 0
    count := 1
    for r := 1; r < len(nums); r++ {
        if nums[r] != nums[l] {
            count++
            l++
            nums[r], nums[l] = nums[l], nums[r]
        }
    }
    return count   
}
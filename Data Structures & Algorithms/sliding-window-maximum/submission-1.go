func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}

    for i:=0; i<len(nums) - k + 1; i ++ {
        maxf := -10001
        for j:= 0; j < k; j ++ {
            maxf = max(maxf, nums[j+i])
        }  
        res = append(res, maxf)
    }
    return res   
}
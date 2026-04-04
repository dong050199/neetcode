func rob(nums []int) int {
    var dfs func(i int) int 
    dfs = func(i int) int {
        if i >= len(nums) {
            return 0
        }

        return max(dfs(i+1), nums[i] + dfs(i + 2))
    }
   
    return dfs(0)
}
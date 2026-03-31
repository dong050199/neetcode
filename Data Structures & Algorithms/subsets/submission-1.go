func subsets(nums []int) [][]int {
    var res [][]int
    dfs(nums, 0, []int{}, &res)
    return res
}

func dfs(nums []int, i int, subset []int, res *[][]int) {
    if i >= len(nums) {
        tmp := make([]int, len(subset))
        copy(tmp, subset)
        *res = append(*res, tmp)
        return  
    }

    subset = append(subset, nums[i])
    dfs(nums, i+1, subset, res)

    subset = subset[:len(subset)-1]
    dfs(nums, i+1, subset, res)
}
func permute(nums []int) [][]int {
	var res [][]int
	backtracking(nums, []int{}, make([]bool, len(nums)), &res)
	return res
}

func backtracking(nums []int, cur []int, pick []bool, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := append([]int{}, cur...)
		*res = append(*res, tmp)
		return
	}

	for i:= 0; i< len(nums); i++ {
		if !pick[i] {
			pick[i] = true
			cur = append(cur, nums[i])
			backtracking(nums, cur, pick, res)
			cur = cur[:len(cur)-1]
			pick[i] = false
			// will back soon after back tracking
		}
	}
}

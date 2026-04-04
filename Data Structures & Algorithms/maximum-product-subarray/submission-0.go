func maxProduct(nums []int) int {
	// brute force solution
	var res int

	for i := 0; i < len(nums); i++ {
		product := nums[i]
		res = max(res, product)
		for j := i + 1; j < len(nums); j ++ {
			product = product * nums[j]
			res = max(res, product)
		}
	}

	return res
}

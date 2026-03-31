func productExceptSelf(nums []int) []int {
	product := 1 
	zero := 0
	for _, num := range nums {
		if num != 0 {
			product = product * num
		} else {
			zero++
		}
	}

	if zero > 1 {
		res := make([]int, len(nums))
		return res
	}

	res := []int{}
	for _, num := range nums {
		if zero == 1 {
			if num == 0 {
				res = append(res, product)
			} else {
				res = append(res, 0)
			}
			continue
		} else {
			res = append(res, product/num)
		}
	}

	return res
}

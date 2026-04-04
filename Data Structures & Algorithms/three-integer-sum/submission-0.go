
func threeSum(nums []int) [][]int {
	// sort array to make it easier when detect duplicate and calculate (later).
	sort.Ints(nums) // O(n log n)

	var res [][]int
	for i, num := range nums {
		if num > 0 {
			break
		}

		if i > 0 && num == nums[i-1] {
			continue
		}

		// two sum to find the sum of elements which have the sum equal to -triples[0]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r]+num > 0 {
				r--
				continue
			}
			if nums[l]+nums[r]+num < 0 {
				l++
				continue
			}
			if nums[l]+nums[r] == -num {
				res = append(res, []int{num, nums[l], nums[r]})
                l ++
                r --
                continue
			}
		}
	}

	return res
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // handle k > len(nums)
	newNums := append(nums[n-k:], nums[:n-k]...)
	copy(nums, newNums)
}
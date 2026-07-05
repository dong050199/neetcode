func countBits(n int) []int {
	ans := make([]int, n+1)
	lastPow2 := 1 // tracks the most recent power of 2 seen
	for i := 1; i <= n; i++ {
		if i == lastPow2*2 {
			lastPow2 = i // update when i itself becomes the next power of 2
		}
		ans[i] = ans[i-lastPow2] + 1 // ans[lastPow2] is always 1, so "+ ans[lastPow2]" == "+ 1"
	}
	return ans
}
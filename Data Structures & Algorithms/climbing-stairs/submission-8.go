func climbStairs(n int) int {
    a , b := 1, 2
	for i := 3; i <= n; i ++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return b
}

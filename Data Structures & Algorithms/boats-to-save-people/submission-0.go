func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l, r := 0, len(people)-1
	res := 0
	for l <= r {
        if l == r {
            res++
            return res
        }
        
		if people[r] + people[l] > limit {
			r--
			res++
			continue
		}

		res++
		r--
		l++
	}
	return res
}
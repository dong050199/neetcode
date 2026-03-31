func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	total := 0
	res := 0

	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			res = i + 1
		}
	}
	return res
}


func sum(req []int) int {
	res := 0
	for _, r := range req {
		res += r
	}
	return res
}
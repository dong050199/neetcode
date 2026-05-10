func canCompleteCircuit(gas []int, cost []int) int {
	totalCost := 0
	totalTank := 0
	for i := range gas {
		totalCost+=cost[i]
		totalTank+=gas[i]
	}

	if totalCost > totalTank {
		return -1
	}

	res := 0
	tank := 0

	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			res = i + 1
		}
	}
	return res
}







func canCompleteCircuit(gas []int, cost []int) int {
    totalGas := 0
	totalCost := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
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

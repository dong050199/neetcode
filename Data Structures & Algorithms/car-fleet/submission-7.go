func carFleet(target int, position []int, speed []int) int {
	carFleet := []float64{}

	timeToPosition := [][]float64{}

	for i := 0; i < len(position); i++ {
		time := float64(target - position[i]) / float64(speed[i])
		timeToPosition = append(timeToPosition, []float64{float64(position[i]), time})
	}

	sort.Slice(timeToPosition, func(i, j int) bool {
    	return timeToPosition[i][0] > timeToPosition[j][0]
	})

	for _, t := range timeToPosition {
		if len(carFleet) == 0 || carFleet[len(carFleet)-1] < t[1] {
			carFleet = append(carFleet, t[1])
		}
	}


	return len(carFleet)
}

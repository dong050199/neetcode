func carFleet(target int, position []int, speed []int) int {
	timeToPos := [][]float64{}
	for i := 0; i < len(position); i++ {
		timeToPos = append(timeToPos, []float64{float64(position[i]), float64(target - position[i])/float64(speed[i])})
	}

	sort.Slice(timeToPos, func(i, j int) bool {
    	return timeToPos[i][0] > timeToPos[j][0]
	})

	carFleet := []float64{}
	for i := 0; i < len(timeToPos); i++ {
		if len(carFleet) > 0 && carFleet[len(carFleet)-1] > timeToPos[i][1] {
			continue
		} else {
			carFleet = append(carFleet, timeToPos[i][1])
		}
	}
	return len(carFleet) - 1
}

func carFleet(target int, position []int, speed []int) int {
	positionTime := [][]float64{}

	for i := 0; i < len(position); i++ {
		time := float64(target - position[i]) / float64(speed[i])
		positionTime = append(positionTime, []float64{float64(position[i]), time})
	}

	sort.Slice(positionTime, func(i, j int) bool {
    	return positionTime[i][0] > positionTime[j][0]
	})

	res := []float64{positionTime[0][1]}
	for i := 1; i < len(positionTime); i++ {
		if len(res) > 0 && positionTime[i][1] > res[len(res)-1] {
			res = append(res, positionTime[i][1])
		}
	}

	return len(res)
}

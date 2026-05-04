func carFleet(target int, position []int, speed []int) int {
	tmp := [][]float64{}
	
	for i := 0; i < len(position); i++ {
		tmp = append(tmp, []float64{float64(position[i]), (float64(target) - float64(position[i])) / float64(speed[i])})
	}

	sort.Slice(tmp, func(i, j int) bool {
    	return tmp[i][0] > tmp[j][0] 
	})

	stack := []float64{}

	for _, t := range tmp {
		if len(stack) == 0 || stack[len(stack)-1] < t[1] {
			stack = append(stack, t[1])
		}
	}
	
	return len(stack)
}
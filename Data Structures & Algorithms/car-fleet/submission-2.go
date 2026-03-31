func carFleet(target int, position []int, speed []int) int {
    pair := make([][2]int, len(position))
    for i:= 0; i < len(position); i++ {
        pair[i] = [2]int{position[i], speed[i]}
    }

	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

    stack := []float64{}
    for i := 0; i < len(position); i++ {
        t := float64((target - pair[i][0]))/float64(pair[i][1])
        if i == 0 {
            stack = append(stack,t)
            continue
        }
        if t > stack[len(stack)-1] {
            stack = append(stack,t)
            continue
        }
    }
    return len(stack)
}

func carFleet(target int, position []int, speed []int) int {
    pair := make([][2]int, len(position))
    for i:= 0; i < len(position); i++ {
        pair = append(pair, [2]int{position[i], speed[i]})
    }

    // sort slice from left to right
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


// this problem can use stack, let think about time when a car come to target from righ to left
// can consider the first car is a car fleet as no cars can be pass this car.
// if the next car have the time to target greater or equal to the first car then it will join with the first car.
// if the next car have time to target less then the first car then it will create new fleet.
// next until reach the target

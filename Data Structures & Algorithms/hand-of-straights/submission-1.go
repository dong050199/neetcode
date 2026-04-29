func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	count := make(map[int]int)
	
	for _, card := range hand {
		count[card]++
	}


	for _, card := range hand {
		if count[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if count[card + i] == 0 {
				return false
			}
			count[card + i]--
		}
	}

	return true
}

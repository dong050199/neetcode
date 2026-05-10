func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {
		return false
	}

	card := make(map[int]int)

	for _, c := range hand {
		card[c]++
	}

    sort.Ints(hand)

	for _, c := range hand {
		if card[c] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if card[c + i] == 0 {
				return false
			}
			card[c+i]--
		}
	}

	return true
}

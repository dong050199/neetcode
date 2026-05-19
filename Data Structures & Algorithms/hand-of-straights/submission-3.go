func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {
		return false
	}

	mp := make(map[int]int)
	for _, c := range hand {
		mp[c]++
	}

	sort.Ints(hand)

	for i := 0; i < len(hand); i++ {
		if mp[hand[i]] == 0 {
			continue
		}


		for j := 0; j < groupSize; j++ {
			if mp[hand[i] + j] < 0 {
				return false
			}
			mp[hand[i] + j]--
		}
	}

    return true
}

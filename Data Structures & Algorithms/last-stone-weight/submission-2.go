func lastStoneWeight(stones []int) int {
	// we need to pick s1 and s2 first (heaviest stones)
	// implement bruteforce 
	if len(stones) <= 1 {
		if len(stones) == 0 {
			return 0
		}
		return stones[0]
	}

	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})

	if s := smash(stones[0], stones[1]); s > 0 {
		stones = stones[2:]
		stones = append(stones, s) 
	} else {
		stones = stones[2:]
	}

	return lastStoneWeight(stones)
}

// return -1 if both s1 and s2 destroyed after smash
func smash(s1, s2 int) int {
	if s1 == s2 {
		return -1
	}

	if s1 > s2 {
		return s1 - s2
	}

	return s2 - s1
}

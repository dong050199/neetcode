func checkInclusion(s1 string, s2 string) bool {
	s1f := make(map[byte]int)

	for i := 0; i < len(s1); i ++ {
		s1f[s1[i]]++
	}


	for i := 0; i < len(s2); i ++ {
		cpS1f := s1f
		if _, exist := cpS1f[s2[i]]; !exist {
			continue
		}

		cpS1f[s2[i]]--
		for j :=  i + 1 ; j < len(s2) ; j ++ {
			if _, exist := cpS1f[s2[j]]; exist {
				cpS1f[s2[j]]--
				continue
			}
			break
		}

		tmp := true
		for _, v := range cpS1f {
			if v != 0 {
				tmp = false
				break
			}
		}

		if tmp {
			return true
		}
	}
	
	return false
}

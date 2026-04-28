func letterCombinations(digits string) []string {
	keyLetterMap := make(map[string][]string)

	keyLetterMap["1"] = []string{}
	keyLetterMap["2"] = []string{"a", "b", "c"}
	keyLetterMap["3"] = []string{"d", "e", "f"}
	keyLetterMap["4"] = []string{"g", "h", "i"}
	keyLetterMap["5"] = []string{"j", "k", "l"}
	keyLetterMap["6"] = []string{"m", "n", "o"}
	keyLetterMap["7"] = []string{"p", "q", "r", "s"}
	keyLetterMap["8"] = []string{"t", "u", "v"}
	keyLetterMap["9"] = []string{"w", "x", "y", "z"}

	dgs := strings.Split(digits, "")
	if len(dgs) == 0 {
		return []string{}
	}

	res := []string{}
	var dfs func(i int, comb string)
	dfs = func(i int, comb string) {
		if i >= len(dgs) {
			res = append(res, comb)
			return
		}

		for _, l := range keyLetterMap[dgs[i]] {
			comb += l
			dfs(i + 1, comb)
			comb = comb[:len(comb)-1]
		}
	}


	dfs(0, "")
	return res
}

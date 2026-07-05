func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		columnNumber--
		last := string(rune('A' + columnNumber%26))
		res = last + res
		columnNumber = columnNumber / 26
	}
    return res
}
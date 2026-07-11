func mergeAlternately(word1 string, word2 string) string {
    res := ""
    i, j := 0, 0
    for  i < len(word1) && j < len(word2) {
        res += string(word1[i]) + string(word2[j])
        i++
        j++
    }

    res += word1[i:]
    res += word2[j:]

    return res
}
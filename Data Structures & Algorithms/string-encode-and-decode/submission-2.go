type Solution struct{
    ln []int
}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        s.ln = append(s.ln, len(str))
        res += str
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    for _, l := range s.ln {
        res = append(res, encoded[:l])
        encoded = encoded[l:]
    }
    return res
}

type Solution struct{}

var delimiter = "$"
func (s *Solution) Encode(strs []string) string {
    str := ""
    for _, s := range strs {
        str += fmt.Sprintf("%d%s%s", len(s), delimiter, s)
    }
    return str
}

func (s *Solution) Decode(encoded string) []string {
    strs := []string{}
    for {
       lengthStr, after, found := strings.Cut(encoded, delimiter)
       if !found {
        break
       }
       encoded = after

       length, err := strconv.Atoi(lengthStr)
       if err != nil {
        break
       }
        str := encoded[:length]
        encoded = encoded[length:]
        strs = append(strs, str)
    }

    return strs
}

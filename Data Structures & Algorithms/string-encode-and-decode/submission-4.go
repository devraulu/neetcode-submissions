type Solution struct{}

var delimiter = "$"

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, s := range strs {
        sb.WriteString(strconv.Itoa(len(s)))
        sb.WriteString(delimiter)
        sb.WriteString(s)
    }

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    i := 0
    res := []string{}
    for i < len(encoded) {
        j := i
        for string(encoded[j]) != delimiter {
            j++
        }
        length, err := strconv.Atoi(encoded[i:j])
        if err != nil {
            break;
        }
        i = j + 1
        res = append(res, encoded[i:i+length])
        i += length
    }

    return res
}

type FrequencyArray [27]int

func groupAnagrams(strs []string) [][]string {
 if len(strs) < 1 {
        return [][]string{}
    }
	var lower func(rune) rune 
	lower = func(c rune) rune {
		if c >= 'A' && c <= 'Z' {
			return c + 'a' - 'A'
		}
		return c
	}

    m := map[FrequencyArray][]string{}
    for _, str := range strs {
		fm := FrequencyArray{}
		for _, r := range str {
			fm[lower(r)-'a']++
		}

        m[fm] = append(m[fm], str)
    }

	arr := [][]string{}
   for _, v := range m {
    arr = append(arr, v)
   }
   return arr
}
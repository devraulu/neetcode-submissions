func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}

	for _, s := range strs {
		var count [26]int
		for	_, r := range s {
			count[r-'a']++
		}
		m[count] = append(m[count], s)
	}

	res := [][]string{}
	for _, group := range m{
		res = append(res, group)
	}

	return res
}

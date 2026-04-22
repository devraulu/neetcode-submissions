
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		r := []rune(str)

		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})

		sorted := string(r)
		m[sorted] = append(m[sorted], str)
	}

	arr := [][]string{}
	for _, v := range m {
		arr = append(arr, v)
	}

	return arr
}

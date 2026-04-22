func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		count := [26]int{}
		for _, r := range str {
			count[r-'a']++ 
		}
		
		m[count] = append(m[count], str)
	}

	result := [][]string{}
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

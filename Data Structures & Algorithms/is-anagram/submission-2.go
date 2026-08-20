func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
	m1 := map[byte]int{}
	m2 := map[byte]int{}

	for i := range len(s) {
		m1[s[i]] = m1[s[i]] + 1
		m2[t[i]] = m2[t[i]] + 1
	}

	for k, v := range m1 {
		if v2 := m2[k]; v2 != v {
			return false
		}
	}

	return true
}

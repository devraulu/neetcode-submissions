func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	fmS := map[rune]int{}
	fmT := map[rune]int{}

	for _, c := range s {
		fmS[c]++
	}
	for _, c := range t {
		fmT[c]++
	}
	for k, v := range fmS {
		if fmT[k] != v {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	fmS := map[rune]int{}
	fmT := map[rune]int{}

	for i, c := range s {
		fmS[c]++
		fmT[rune(t[i])]++
	}

	for k, v := range fmS {
		if fmT[k] != v {
			return false
		}
	}
	return true
}

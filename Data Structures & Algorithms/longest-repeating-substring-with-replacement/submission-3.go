func characterReplacement(s string, k int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		fm := map[byte]int{}
		maxf := 0
		for j := i; j < len(s); j++ {
			fm[s[j]]++
			maxf = max(maxf, fm[s[j]])

			if j - i + 1 - maxf <= k {
				res = max(res, j - i + 1)
			}
		}
	}

	return res
}

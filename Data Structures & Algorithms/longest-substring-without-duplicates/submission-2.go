func lengthOfLongestSubstring(s string) int {
	set := map[byte]bool{}
	l, longest := 0, 0

	for r := 0; r < len(s); r++ {
		for set[s[r]] {
			delete(set, s[l])
			l++
		}

		set[s[r]] = true
		if r - l + 1 > longest {
			longest = r - l + 1
		}

	}

	return longest
}

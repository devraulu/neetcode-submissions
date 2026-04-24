func lengthOfLongestSubstring(s string) int {
	mp := map[byte]int{}
	l, longest := 0, 0

	for r := 0; r < len(s); r++ {
		if i, ok := mp[s[r]]; ok {
			// max because if i < l then "duplicate" is outside of window
			l = max(i + 1, l)
		}
		mp[s[r]] = r
		
		if r - l + 1 > longest {
			longest = r - l + 1
		}
	}

	return longest
}
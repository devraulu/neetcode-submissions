func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, n := range nums {
		set[n] = true
	}

	max := 0
	for _, n := range nums {
		if !set[n - 1] {
			length := 1
			for set[n+length] {
				length++
			}
			if length > max {
				max = length
			}

		}
	}
	return max
}

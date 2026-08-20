func topKFrequent(nums []int, k int) []int {
	fm := map[int]int{}
	for _, n := range nums {
		fm[n]++
	}

	n := len(nums)+1
	fb := make([][]int, n)
	for k, v := range fm {
		fb[v] = append(fb[v], k)
	}

	res := []int{}
	count := 0
	for i := len(fb) - 1; i >= 0; i-- {
        b := fb[i]
		for _, n := range b {
			res = append(res, n)
			count++
			if count == k {
				return res
			}
		}
	}

	return []int{}
}

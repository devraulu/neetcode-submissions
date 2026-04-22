func topKFrequent(nums []int, k int) []int {
	fm := map[int]int{}
	for _, n := range nums {
		fm[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for k, v := range fm {
		bucket[v] = append(bucket[v], k)
	}

	top := []int{}

	for i := len(bucket)-1; i > 0; i-- {
		for _, n := range bucket[i] {
			top = append(top, n)
			if len(top) == k {
				return top
			}
		}
	}

	return top 
}

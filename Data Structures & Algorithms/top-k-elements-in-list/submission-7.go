func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums) + 1)
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}

	result := []int{}

	for i := len(bucket) - 1; i > 0; i-- {
		for _, n := range bucket[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}

	return []int{}
}

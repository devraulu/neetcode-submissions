func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		diff := target - n
		m[diff] = i
	}

	for i, n := range nums {
		if j, ok := m[n]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

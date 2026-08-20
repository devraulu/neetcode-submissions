func twoSum(nums []int, target int) []int {
   	m := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if j, ok := m[diff]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
    return []int{} 
}

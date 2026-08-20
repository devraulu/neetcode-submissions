func twoSum(nums []int, target int) []int {
   	m := map[int]int{}
	for i, n := range nums {
		if j, ok := m[n]; ok {
			if j > i{
				return []int{i, j}
			} else {
				return []int{j, i}
			}
		}
		diff := target - n
		m[diff] = i
	}
    return []int{} 
}

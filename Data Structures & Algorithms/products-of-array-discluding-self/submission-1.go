func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	pre, suf := 1, 1
	for i := 0; i < n; i++ {
		x := nums[i]
		res[i] = pre 
		pre = pre * x
	}
	for i := 0; i < n; i++ {
		y := nums[n - i - 1]
		res[n - i - 1] = res[n - i - 1] * suf
		suf = suf * y
	}
	return res
}

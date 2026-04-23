func maxProfit(prices []int) int {
	max := 0
	l := 0
	r := 1

	for r < len(prices) {
		if prices[r] > prices[l] {
			profit := prices[r] - prices[l]
			if profit > max {
				max = profit
			}
		} else {
			l = r
		}
		r++
	}

	return max
}

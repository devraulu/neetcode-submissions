func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]struct{}, n)
	cols := make([]map[byte]struct{}, n)
	subgrids := make([]map[byte]struct{}, n)
	for i := range n {
		rows[i] = map[byte]struct{}{}
		cols[i] = map[byte]struct{}{}
		subgrids[i] = map[byte]struct{}{}
	}

	for r := range n {
		for c := range n {
			cell := board[r][c]
			if cell == '.' {
				continue
			}

			subgrid := int((math.Floor(float64(r) / 3) * 3) + math.Floor(float64(c) / 3))
			if _, ok := rows[r][cell]; ok {
				return false
			}
			if _, ok := cols[c][cell]; ok {
				return false
			}
			if _, ok := subgrids[subgrid][cell]; ok {
				return false
			}
			rows[r][cell] = struct{}{}
			cols[c][cell] = struct{}{}
			subgrids[subgrid][cell] = struct{}{}
		}
	}
	return true
}

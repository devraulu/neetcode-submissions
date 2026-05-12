func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)
    
	var dfs func(r, c int, vis map[[2]int]bool, prev int)
	dfs = func(r, c int, vis map[[2]int]bool, prev int) {
		coord := [2]int{r, c}
		if vis[coord] || r < 0 || c < 0 || r == rows || c == cols || heights[r][c] < prev {
			return
		}

		vis[coord] = true

		dfs(r+1, c, vis, heights[r][c])
		dfs(r-1, c, vis, heights[r][c])
		dfs(r, c+1, vis, heights[r][c])
		dfs(r, c-1, vis, heights[r][c])
	}

	for c := range cols {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, atl, heights[rows-1][c])
	}

	for r := range rows {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, atl, heights[r][cols-1])
	}

	result := make([][]int, 0)
	for r := range rows {
		for c := range cols {
			coord := [2]int{r,c}
			if pac[coord] && atl[coord] {
				result = append(result, []int{r,c})
			}
		}
	}

	return result
}

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var maxa int
	vis := map[node]bool{}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 && !vis[node{r,c}] {
				count := BFS(r, c, grid, &vis)
				fmt.Println(count)
				maxa = max(maxa, count)
			}
		}
	}

	return maxa
}

func BFS(r, c int, grid [][]int, vis *map[node]bool) int {
	rows, cols := len(grid), len(grid[0])
	q := []node{}
	q = append(q, node{r,c})

	(*vis)[node{r,c}] = true
	dirs := []node{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	count := 1

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		r, c := curr.X, curr.Y
		for _, dir := range dirs {
			r, c := r+dir.X, c+dir.Y
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != 1 || (*vis)[node{r,c}] {
				continue
			}
			q = append(q, node{r,c})
			(*vis)[node{r,c}] = true
			count++
		}
	}
	return count
}

type node struct {
	X int
	Y int
}
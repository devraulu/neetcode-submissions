func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	islands := 0
	vis := map[Node]bool{}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' && !vis[Node{r, c}] {
				BFS(r, c, grid, &vis)
				islands++
			}
		}
	}

	return islands
}

func BFS(r, c int, grid [][]byte, vis *map[Node]bool) {
	q := []Node{}
	q = append(q, Node{r, c})
	(*vis)[Node{r, c}] = true

	dirs := []Node{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	rows, cols := len(grid), len(grid[0])

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		r, c = curr.X, curr.Y
		for _, dir := range dirs {
			nr := r + dir.X
			nc := c + dir.Y
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != '1' || (*vis)[Node{nr, nc}] {
				continue
			}
			q = append(q, Node{nr, nc})
			(*vis)[Node{nr, nc}] = true
		}
	}

}

type Node struct {
	X int
	Y int
}

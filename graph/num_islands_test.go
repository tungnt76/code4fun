// https://leetcode.com/problems/number-of-islands/
package graph

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	var m1DFS func(r, c int)
	m1DFS = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		m1DFS(r-1, c)
		m1DFS(r+1, c)
		m1DFS(r, c+1)
		m1DFS(r, c-1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				m1DFS(i, j)
			}
		}
	}

	return count
}


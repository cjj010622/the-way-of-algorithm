package day43

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

// dfs
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' {
			return
		}

		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	var res int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				res++
				dfs(row, col)
			}
		}
	}
	return res
}

// bfs
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	res := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				res++

				grid[row][col] = '0'
				queue := [][]int{{row, col}}

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]

					r, c := cur[0], cur[1]
					for _, direction := range directions {
						newR, newC := r+direction[0], c+direction[1]
						if newR < rows && newC < cols && newR >= 0 && newC >= 0 && grid[newR][newC] == '1' {
							grid[newR][newC] = '0'
							queue = append(queue, []int{newR, newC})
						}
					}
				}
			}
		}
	}
	return res
}

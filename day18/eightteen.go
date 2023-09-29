package day18

/*
骑士在一张 n x n 的棋盘上巡视。在 有效 的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。

给你一个 n x n 的整数矩阵 grid ，由范围 [0, n * n - 1] 内的不同整数组成，其中 grid[row][col] 表示单元格 (row, col) 是骑士访问的第 grid[row][col] 个单元格。骑士的行动是从下标 0 开始的。

如果 grid 表示了骑士的有效巡视方案，返回 true；否则返回 false。

注意，骑士行动时可以垂直移动两个格子且水平移动一个格子，或水平移动两个格子且垂直移动一个格子。下图展示了骑士从某个格子出发可能的八种行动路线。
*/

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}

	type pair struct {
		x, y int
	}

	n := len(grid)
	pos := make([]pair, n*n)

	for i, row := range grid {
		for j, dst := range row {
			pos[dst] = pair{i, j}
		}
	}

	for i := 1; i < n*n; i++ {
		dx := abs(pos[i].x - pos[i-1].x)
		dy := abs(pos[i].y - pos[i-1].y)
		ok := (dx == 1 && dy == 2) || (dx == 2 && dy == 1)
		if !ok {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

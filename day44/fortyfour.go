package day44

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= rows || col >= cols || board[row][col] != 'O' {
			return
		}

		board[row][col] = '#'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	// 处理边界上的 'O'
	for row := 0; row < rows; row++ {
		dfs(row, 0)
		dfs(row, cols-1)
	}
	for col := 0; col < cols; col++ {
		dfs(0, col)
		dfs(rows-1, col)
	}

	// 将剩余的 'O' 替换回来，将 '#' 恢复为 'O'
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == '#' {
				board[row][col] = 'O'
			}
		}
	}
}

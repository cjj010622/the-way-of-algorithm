package day21

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		boxes[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			if rows[i][num] {
				return false
			}
			rows[i][num] = true

			if cols[j][num] {
				return false
			}
			cols[j][num] = true

			boxIndex := (i/3)*3 + j/3
			if boxes[boxIndex][num] {
				return false
			}
			boxes[boxIndex][num] = true
		}
	}
	return true
}

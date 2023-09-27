package day22

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	rowBegin, rowEnd := 0, m-1
	colBegin, colEnd := 0, n-1
	result := make([]int, 0)
	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			result = append(result, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[colEnd][i])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i++ {
				result = append(result, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i <= rowBegin; i++ {
				result = append(result, matrix[i][colBegin])
			}
		}
		colBegin++
	}
	return result
}

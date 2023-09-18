package day9

/*
在一个 8x8 的棋盘上，放置着若干「黑皇后」和一个「白国王」。

给定一个由整数坐标组成的数组 queens ，表示黑皇后的位置；以及一对坐标 king ，表示白国王的位置，返回所有可以攻击国王的皇后的坐标(任意顺序)。
*/

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	n := 8
	s := [8][8]bool{}
	for _, q := range queens {
		s[q[0]][q[1]] = true
	}

	var ans [][]int
	for a := -1; a < 2; a++ {
		for b := -1; b < 2; b++ {
			if a == 0 && b == 0 {
				continue
			}

			x, y := king[0]+a, king[1]+b
			for 0 <= x && x < n && 0 <= y && y < n {
				if s[x][y] {
					ans = append(ans, []int{x, y})
					break
				}
				x += a
				y += b
			}
		}
	}
	return ans
}
